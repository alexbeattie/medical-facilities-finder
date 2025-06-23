<template>
  <div
    v-if="isVisible"
    class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4"
    @click="closeModal"
  >
    <div
      class="bg-white rounded-lg shadow-xl max-w-4xl w-full max-h-[90vh] overflow-hidden"
      @click.stop
    >
      <div
        class="flex items-center justify-between p-6 border-b border-gray-200"
      >
        <div>
          <h3 class="text-xl font-semibold text-gray-900">
            🏥 Submit New ABA Center
          </h3>
          <p class="text-sm text-gray-500">
            Help us expand our database by submitting a new ABA therapy center
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

      <div class="overflow-y-auto max-h-[calc(90vh-120px)]">
        <form @submit.prevent="submitForm" class="p-6 space-y-6">
          <!-- Basic Information -->
          <div class="bg-blue-50 p-4 rounded-lg">
            <h4 class="text-lg font-medium text-blue-900 mb-4">
              📋 Basic Information
            </h4>
            <div class="grid grid-cols-1 lg:grid-cols-2 gap-4">
              <div class="lg:col-span-2">
                <label
                  for="name"
                  class="block text-sm font-medium text-gray-700 mb-1"
                  >Center Name *</label
                >
                <input
                  id="name"
                  v-model="form.name"
                  type="text"
                  required
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                  placeholder="Enter the full name of the ABA center"
                />
              </div>
              <div>
                <label
                  for="service_type"
                  class="block text-sm font-medium text-gray-700 mb-1"
                  >Service Type *</label
                >
                <select
                  id="service_type"
                  v-model="form.service_type"
                  required
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                >
                  <option value="">Select service type</option>
                  <option value="ABA Therapy">ABA Therapy</option>
                  <option value="Early Intervention">Early Intervention</option>
                  <option value="School-Based Services">
                    School-Based Services
                  </option>
                  <option value="Adult Services">Adult Services</option>
                </select>
              </div>
              <div>
                <label
                  for="phone"
                  class="block text-sm font-medium text-gray-700 mb-1"
                  >Phone Number *</label
                >
                <input
                  id="phone"
                  v-model="form.phone"
                  type="tel"
                  required
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                  placeholder="(555) 123-4567"
                />
              </div>
            </div>
          </div>

          <!-- Location Information -->
          <div class="bg-green-50 p-4 rounded-lg">
            <h4 class="text-lg font-medium text-green-900 mb-4">
              📍 Location Information
            </h4>
            <div class="grid grid-cols-1 lg:grid-cols-2 gap-4">
              <div class="lg:col-span-2">
                <label
                  for="address"
                  class="block text-sm font-medium text-gray-700 mb-1"
                  >Street Address *</label
                >
                <input
                  id="address"
                  v-model="form.address"
                  type="text"
                  required
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                  placeholder="123 Main Street, Suite 100"
                />
              </div>
              <div>
                <label
                  for="city"
                  class="block text-sm font-medium text-gray-700 mb-1"
                  >City *</label
                >
                <input
                  id="city"
                  v-model="form.city"
                  type="text"
                  required
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                  placeholder="Los Angeles"
                />
              </div>
              <div>
                <label
                  for="state"
                  class="block text-sm font-medium text-gray-700 mb-1"
                  >State *</label
                >
                <select
                  id="state"
                  v-model="form.state"
                  required
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                >
                  <option value="">Select state</option>
                  <option value="CA">California</option>
                  <option value="AZ">Arizona</option>
                  <option value="NV">Nevada</option>
                  <option value="OR">Oregon</option>
                  <option value="WA">Washington</option>
                </select>
              </div>
              <div>
                <label
                  for="zip_code"
                  class="block text-sm font-medium text-gray-700 mb-1"
                  >ZIP Code *</label
                >
                <input
                  id="zip_code"
                  v-model="form.zip_code"
                  type="text"
                  required
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                  placeholder="90210"
                />
              </div>
            </div>
          </div>

          <!-- Submitter Information -->
          <div class="bg-orange-50 p-4 rounded-lg">
            <h4 class="text-lg font-medium text-orange-900 mb-4">
              👤 Your Information
            </h4>
            <div class="grid grid-cols-1 lg:grid-cols-2 gap-4">
              <div>
                <label
                  for="submitter_name"
                  class="block text-sm font-medium text-gray-700 mb-1"
                  >Your Name *</label
                >
                <input
                  id="submitter_name"
                  v-model="form.submitter_name"
                  type="text"
                  required
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                  placeholder="Your full name"
                />
              </div>
              <div>
                <label
                  for="submitter_email"
                  class="block text-sm font-medium text-gray-700 mb-1"
                  >Your Email *</label
                >
                <input
                  id="submitter_email"
                  v-model="form.submitter_email"
                  type="email"
                  required
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                  placeholder="your.email@example.com"
                />
              </div>
              <div class="lg:col-span-2">
                <label
                  for="relationship"
                  class="block text-sm font-medium text-gray-700 mb-1"
                  >Relationship to Center *</label
                >
                <select
                  id="relationship"
                  v-model="form.relationship"
                  required
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                >
                  <option value="">Select your relationship</option>
                  <option value="Staff Member">Staff Member</option>
                  <option value="Administrator">Administrator/Director</option>
                  <option value="Parent">Parent</option>
                  <option value="Family Member">Family Member</option>
                  <option value="Healthcare Provider">
                    Healthcare Provider
                  </option>
                  <option value="Service Provider">Service Provider</option>
                  <option value="Other">Other</option>
                </select>
              </div>
            </div>
          </div>

          <!-- Error/Success Messages -->
          <div
            v-if="error"
            class="p-4 bg-red-50 border border-red-200 rounded-lg"
          >
            <div class="flex items-center text-red-800">
              <span class="text-lg mr-2">⚠️</span>
              <div>
                <div class="font-medium">Submission Error</div>
                <div class="text-sm">{{ error }}</div>
              </div>
            </div>
          </div>

          <div
            v-if="submitted && !error"
            class="p-4 bg-green-50 border border-green-200 rounded-lg"
          >
            <div class="flex items-center text-green-800">
              <span class="text-lg mr-2">✅</span>
              <div>
                <div class="font-medium">Submission Successful!</div>
                <div class="text-sm">
                  Thank you for contributing to our ABA center directory.
                </div>
              </div>
            </div>
          </div>

          <!-- Actions -->
          <div class="flex space-x-3 pt-4 border-t border-gray-200">
            <button
              type="button"
              @click="closeModal"
              class="flex-1 bg-gray-200 text-gray-800 py-3 px-4 rounded-lg hover:bg-gray-300 transition-colors duration-200"
            >
              Cancel
            </button>
            <button
              type="submit"
              :disabled="loading || submitted"
              class="flex-1 bg-blue-600 text-white py-3 px-4 rounded-lg hover:bg-blue-700 disabled:bg-gray-300 disabled:cursor-not-allowed transition-colors duration-200 flex items-center justify-center space-x-2"
            >
              <span v-if="loading" class="inline-block animate-spin">⟳</span>
              <span v-else-if="submitted">✅</span>
              <span v-else>🏥</span>
              <span>{{
                loading
                  ? "Submitting..."
                  : submitted
                  ? "Submitted"
                  : "Submit ABA Center"
              }}</span>
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, reactive } from "vue";
import { apiService } from "@/services/api";

export default {
  name: "ABASubmissionForm",
  props: {
    isVisible: {
      type: Boolean,
      default: false,
    },
  },
  emits: ["close", "submitted"],
  setup(props, { emit }) {
    const form = reactive({
      name: "",
      service_type: "",
      phone: "",
      address: "",
      city: "",
      state: "",
      zip_code: "",
      submitter_name: "",
      submitter_email: "",
      relationship: "",
    });

    const loading = ref(false);
    const error = ref("");
    const submitted = ref(false);

    const resetForm = () => {
      Object.assign(form, {
        name: "",
        service_type: "",
        phone: "",
        address: "",
        city: "",
        state: "",
        zip_code: "",
        submitter_name: "",
        submitter_email: "",
        relationship: "",
      });
      error.value = "";
      submitted.value = false;
    };

    const closeModal = () => {
      if (!loading.value) {
        emit("close");
        setTimeout(resetForm, 300);
      }
    };

    const submitForm = async () => {
      loading.value = true;
      error.value = "";

      try {
        const submissionData = {
          ...form,
          submitted_at: new Date().toISOString(),
        };

        await apiService.submitABACenter(submissionData);
        submitted.value = true;
        emit("submitted", submissionData);

        setTimeout(() => {
          if (submitted.value) {
            closeModal();
          }
        }, 3000);
      } catch (err) {
        console.error("Error submitting ABA center:", err);
        error.value =
          err.message ||
          "Failed to submit ABA center information. Please try again.";
      } finally {
        loading.value = false;
      }
    };

    return {
      form,
      loading,
      error,
      submitted,
      closeModal,
      submitForm,
      resetForm,
    };
  },
};
</script>

<style scoped>
.z-50 {
  z-index: 50;
}
</style>
