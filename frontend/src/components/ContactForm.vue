<template>
  <!-- Overlay -->
  <div
    v-if="isVisible"
    class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4"
    @click="closeModal"
  >
    <!-- Modal -->
    <div
      class="bg-white rounded-lg shadow-xl max-w-md w-full max-h-[90vh] overflow-hidden"
      @click.stop
    >
      <!-- Header -->
      <div
        class="flex items-center justify-between p-6 border-b border-gray-200"
      >
        <div>
          <h3 class="text-lg font-semibold text-gray-900">
            Contact {{ facilityName }}
          </h3>
          <p class="text-sm text-gray-500">
            Send a message to inquire about services
          </p>
        </div>
        <button
          @click="closeModal"
          class="text-gray-400 hover:text-gray-600 transition-colors"
        >
          <svg
            class="w-6 h-6"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M6 18L18 6M6 6l12 12"
            ></path>
          </svg>
        </button>
      </div>

      <!-- Form -->
      <form @submit.prevent="submitForm" class="p-6 space-y-4">
        <!-- Name -->
        <div>
          <label
            for="name"
            class="block text-sm font-medium text-gray-700 mb-1"
          >
            Full Name *
          </label>
          <input
            id="name"
            v-model="form.name"
            type="text"
            required
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            placeholder="Enter your full name"
          />
        </div>

        <!-- Email -->
        <div>
          <label
            for="email"
            class="block text-sm font-medium text-gray-700 mb-1"
          >
            Email Address *
          </label>
          <input
            id="email"
            v-model="form.email"
            type="email"
            required
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            placeholder="Enter your email address"
          />
        </div>

        <!-- Phone -->
        <div>
          <label
            for="phone"
            class="block text-sm font-medium text-gray-700 mb-1"
          >
            Phone Number
          </label>
          <input
            id="phone"
            v-model="form.phone"
            type="tel"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            placeholder="Enter your phone number"
          />
        </div>

        <!-- Inquiry Type -->
        <div>
          <label
            for="inquiry_type"
            class="block text-sm font-medium text-gray-700 mb-1"
          >
            Inquiry Type *
          </label>
          <select
            id="inquiry_type"
            v-model="form.inquiry_type"
            required
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          >
            <option value="">Select an inquiry type</option>
            <option value="general_info">General Information</option>
            <option value="services">Services Inquiry</option>
            <option value="insurance">Insurance Questions</option>
            <option value="appointment">Schedule Appointment</option>
            <option value="waitlist">Waitlist Information</option>
            <option value="support">Support Services</option>
            <option value="other">Other</option>
          </select>
        </div>

        <!-- Age of Individual (for ABA centers) -->
        <div v-if="facilityType === 'aba_center'">
          <label for="age" class="block text-sm font-medium text-gray-700 mb-1">
            Age of Individual Seeking Services
          </label>
          <input
            id="age"
            v-model.number="form.age"
            type="number"
            min="0"
            max="100"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            placeholder="Enter age"
          />
        </div>

        <!-- Diagnosis (for resource centers) -->
        <div
          v-if="
            facilityType === 'resource_center' || facilityType === 'aba_center'
          "
        >
          <label
            for="diagnosis"
            class="block text-sm font-medium text-gray-700 mb-1"
          >
            Diagnosis (if applicable)
          </label>
          <select
            id="diagnosis"
            v-model="form.diagnosis"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          >
            <option value="">Select a diagnosis</option>
            <option value="Autism">Autism Spectrum Disorder</option>
            <option value="ADHD">ADHD</option>
            <option value="Anxiety">Anxiety</option>
            <option value="Depression">Depression</option>
            <option value="Learning Disability">Learning Disability</option>
            <option value="Other">Other</option>
          </select>
        </div>

        <!-- Insurance Provider -->
        <div v-if="facilityType === 'aba_center'">
          <label
            for="insurance"
            class="block text-sm font-medium text-gray-700 mb-1"
          >
            Insurance Provider
          </label>
          <input
            id="insurance"
            v-model="form.insurance"
            type="text"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            placeholder="Enter your insurance provider"
          />
        </div>

        <!-- Message -->
        <div>
          <label
            for="message"
            class="block text-sm font-medium text-gray-700 mb-1"
          >
            Message *
          </label>
          <textarea
            id="message"
            v-model="form.message"
            required
            rows="4"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent resize-none"
            placeholder="Please describe your inquiry in detail..."
          ></textarea>
        </div>

        <!-- Preferred Contact Method -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            Preferred Contact Method
          </label>
          <div class="flex space-x-4">
            <label class="flex items-center">
              <input
                v-model="form.preferred_contact"
                type="radio"
                value="email"
                class="text-blue-600 focus:ring-blue-500"
              />
              <span class="ml-2 text-sm text-gray-700">Email</span>
            </label>
            <label class="flex items-center">
              <input
                v-model="form.preferred_contact"
                type="radio"
                value="phone"
                class="text-blue-600 focus:ring-blue-500"
              />
              <span class="ml-2 text-sm text-gray-700">Phone</span>
            </label>
          </div>
        </div>

        <!-- Consent -->
        <div>
          <label class="flex items-start space-x-2">
            <input
              v-model="form.consent"
              type="checkbox"
              required
              class="mt-1 rounded text-blue-600 focus:ring-blue-500"
            />
            <span class="text-sm text-gray-700">
              I consent to being contacted by this facility regarding my
              inquiry. *
            </span>
          </label>
        </div>

        <!-- Error Message -->
        <div
          v-if="error"
          class="p-3 bg-red-50 border border-red-200 rounded-lg"
        >
          <div class="flex items-center text-red-800">
            <span class="text-lg mr-2">⚠️</span>
            <div>
              <div class="font-medium">Error</div>
              <div class="text-sm">{{ error }}</div>
            </div>
          </div>
        </div>

        <!-- Success Message -->
        <div
          v-if="submitted && !error"
          class="p-3 bg-green-50 border border-green-200 rounded-lg"
        >
          <div class="flex items-center text-green-800">
            <span class="text-lg mr-2">✅</span>
            <div>
              <div class="font-medium">Message Sent!</div>
              <div class="text-sm">
                Your inquiry has been submitted successfully.
              </div>
            </div>
          </div>
        </div>

        <!-- Actions -->
        <div class="flex space-x-3 pt-4">
          <button
            type="button"
            @click="closeModal"
            class="flex-1 bg-gray-200 text-gray-800 py-2 px-4 rounded-lg hover:bg-gray-300 transition-colors duration-200"
          >
            Cancel
          </button>
          <button
            type="submit"
            :disabled="loading || submitted"
            class="flex-1 bg-blue-600 text-white py-2 px-4 rounded-lg hover:bg-blue-700 disabled:bg-gray-300 disabled:cursor-not-allowed transition-colors duration-200 flex items-center justify-center space-x-2"
          >
            <span v-if="loading" class="inline-block animate-spin">⟳</span>
            <span v-else-if="submitted">✅</span>
            <span v-else>📧</span>
            <span>{{
              loading ? "Sending..." : submitted ? "Sent" : "Send Message"
            }}</span>
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
import { ref, reactive, computed, watch } from "vue";
import { useFacilities } from "@/composables/useFacilities";

export default {
  name: "ContactForm",
  props: {
    facility: {
      type: Object,
      default: null,
    },
    facilityType: {
      type: String,
      default: "facility",
    },
    isVisible: {
      type: Boolean,
      default: false,
    },
  },
  emits: ["close", "submitted"],
  setup(props, { emit }) {
    const { submitContactForm, loading, error } = useFacilities();

    // Form state
    const form = reactive({
      name: "",
      email: "",
      phone: "",
      inquiry_type: "",
      message: "",
      age: null,
      diagnosis: "",
      insurance: "",
      preferred_contact: "email",
      consent: false,
    });

    const submitted = ref(false);

    // Computed
    const facilityName = computed(() => {
      return (
        props.facility?.name || props.facility?.center_name || "this facility"
      );
    });

    // Watch for visibility changes to reset form
    watch(
      () => props.isVisible,
      (newValue) => {
        if (newValue) {
          resetForm();
        }
      }
    );

    // Methods
    const resetForm = () => {
      Object.assign(form, {
        name: "",
        email: "",
        phone: "",
        inquiry_type: "",
        message: "",
        age: null,
        diagnosis: "",
        insurance: "",
        preferred_contact: "email",
        consent: false,
      });
      submitted.value = false;
    };

    const closeModal = () => {
      if (!loading.value) {
        emit("close");
      }
    };

    const submitForm = async () => {
      try {
        const formData = {
          facility_id: props.facility?.id,
          facility_type: props.facilityType,
          facility_name: facilityName.value,
          ...form,
          submitted_at: new Date().toISOString(),
        };

        await submitContactForm(formData);
        submitted.value = true;
        emit("submitted", formData);

        // Auto-close after success
        setTimeout(() => {
          if (submitted.value) {
            closeModal();
          }
        }, 2000);
      } catch (err) {
        console.error("Error submitting contact form:", err);
      }
    };

    return {
      // Data
      form,
      submitted,
      loading,
      error,

      // Computed
      facilityName,

      // Methods
      closeModal,
      submitForm,
      resetForm,
    };
  },
};
</script>

<style scoped>
/* Ensure modal appears above everything else */
.z-50 {
  z-index: 50;
}

/* Custom radio button styling */
input[type="radio"] {
  color: #3b82f6;
}

input[type="radio"]:focus {
  ring-color: #3b82f6;
}

/* Custom checkbox styling */
input[type="checkbox"] {
  color: #3b82f6;
}

input[type="checkbox"]:focus {
  ring-color: #3b82f6;
}
</style>
