<!-- Regional Center Update Form -->
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
            📋 Update Regional Center Information
          </h3>
          <p class="text-sm text-gray-500">
            Help us keep regional center information accurate and up-to-date
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
          <!-- Regional Center Selection -->
          <div class="bg-purple-50 p-4 rounded-lg">
            <h4 class="text-lg font-medium text-purple-900 mb-4">
              🏢 Select Regional Center
            </h4>
            <div class="grid grid-cols-1 gap-4">
              <div>
                <label
                  for="regional_center_name"
                  class="block text-sm font-medium text-gray-700 mb-1"
                  >Regional Center *</label
                >
                <select
                  id="regional_center_name"
                  v-model="form.regional_center_name"
                  required
                  @change="onRegionalCenterChange"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                >
                  <option value="">Select a regional center</option>
                  <option
                    v-for="center in regionalCenters"
                    :key="center.id"
                    :value="center.regional_center"
                  >
                    {{ center.regional_center }}
                  </option>
                  <option value="other">🔹 Other (Please specify)</option>
                </select>
              </div>

              <!-- Other regional center name -->
              <div v-if="form.regional_center_name === 'other'">
                <label
                  for="other_center_name"
                  class="block text-sm font-medium text-gray-700 mb-1"
                  >Regional Center Name *</label
                >
                <input
                  id="other_center_name"
                  v-model="form.other_center_name"
                  type="text"
                  required
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                  placeholder="Enter the name of the regional center"
                />
              </div>
            </div>
          </div>

          <!-- Update Type -->
          <div class="bg-blue-50 p-4 rounded-lg">
            <h4 class="text-lg font-medium text-blue-900 mb-4">
              📝 Type of Update
            </h4>
            <div>
              <label
                for="update_type"
                class="block text-sm font-medium text-gray-700 mb-1"
                >What type of information are you updating? *</label
              >
              <select
                id="update_type"
                v-model="form.update_type"
                required
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              >
                <option value="">Select update type</option>
                <option value="contact_info">Contact Information</option>
                <option value="services">Services & Programs</option>
                <option value="hours">Operating Hours</option>
                <option value="eligibility">Eligibility Criteria</option>
                <option value="accessibility">Accessibility Information</option>
                <option value="staff">Staff Changes</option>
                <option value="location">Location/Address Changes</option>
                <option value="other">Other</option>
              </select>
            </div>
          </div>

          <!-- Update Description -->
          <div class="bg-green-50 p-4 rounded-lg">
            <h4 class="text-lg font-medium text-green-900 mb-4">
              📄 Update Description
            </h4>
            <div>
              <label
                for="update_description"
                class="block text-sm font-medium text-gray-700 mb-1"
              >
                Describe the update or additional information *
              </label>
              <textarea
                id="update_description"
                v-model="form.update_description"
                required
                rows="4"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="Please provide detailed information about what should be updated..."
              ></textarea>
            </div>
          </div>

          <!-- Contact Information Updates -->
          <div
            v-if="form.update_type === 'contact_info'"
            class="bg-orange-50 p-4 rounded-lg"
          >
            <h4 class="text-lg font-medium text-orange-900 mb-4">
              📞 Contact Information
            </h4>
            <div class="grid grid-cols-1 lg:grid-cols-2 gap-4">
              <div>
                <label
                  for="new_phone"
                  class="block text-sm font-medium text-gray-700 mb-1"
                  >New Phone Number</label
                >
                <input
                  id="new_phone"
                  v-model="form.new_phone"
                  type="tel"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                  placeholder="(555) 123-4567"
                />
              </div>
              <div>
                <label
                  for="new_email"
                  class="block text-sm font-medium text-gray-700 mb-1"
                  >New Email Address</label
                >
                <input
                  id="new_email"
                  v-model="form.new_email"
                  type="email"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                  placeholder="contact@regionalcenter.org"
                />
              </div>
              <div class="lg:col-span-2">
                <label
                  for="new_website"
                  class="block text-sm font-medium text-gray-700 mb-1"
                  >New Website</label
                >
                <input
                  id="new_website"
                  v-model="form.new_website"
                  type="url"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                  placeholder="https://www.regionalcenter.org"
                />
              </div>
              <div class="lg:col-span-2">
                <label
                  for="new_address"
                  class="block text-sm font-medium text-gray-700 mb-1"
                  >New Address</label
                >
                <input
                  id="new_address"
                  v-model="form.new_address"
                  type="text"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                  placeholder="123 Main Street, Suite 100, Los Angeles, CA 90210"
                />
              </div>
            </div>
          </div>

          <!-- Services Information -->
          <div
            v-if="form.update_type === 'services'"
            class="bg-indigo-50 p-4 rounded-lg"
          >
            <h4 class="text-lg font-medium text-indigo-900 mb-4">
              🛠️ Services & Programs
            </h4>
            <div class="space-y-4">
              <div>
                <label
                  for="services_offered"
                  class="block text-sm font-medium text-gray-700 mb-1"
                  >Services Offered</label
                >
                <textarea
                  id="services_offered"
                  v-model="form.services_offered"
                  rows="3"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                  placeholder="List the services and programs offered..."
                ></textarea>
              </div>
              <div>
                <label
                  for="special_programs"
                  class="block text-sm font-medium text-gray-700 mb-1"
                  >Special Programs</label
                >
                <textarea
                  id="special_programs"
                  v-model="form.special_programs"
                  rows="3"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                  placeholder="Describe any special programs or initiatives..."
                ></textarea>
              </div>
            </div>
          </div>

          <!-- Operating Hours -->
          <div
            v-if="form.update_type === 'hours'"
            class="bg-yellow-50 p-4 rounded-lg"
          >
            <h4 class="text-lg font-medium text-yellow-900 mb-4">
              🕒 Operating Hours
            </h4>
            <div>
              <label
                for="operating_hours"
                class="block text-sm font-medium text-gray-700 mb-1"
                >Operating Hours</label
              >
              <textarea
                id="operating_hours"
                v-model="form.operating_hours"
                rows="4"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="Monday - Friday: 8:00 AM - 5:00 PM&#10;Saturday: 9:00 AM - 1:00 PM&#10;Sunday: Closed&#10;&#10;Include any special hours or closures..."
              ></textarea>
            </div>
          </div>

          <!-- Eligibility Criteria -->
          <div
            v-if="form.update_type === 'eligibility'"
            class="bg-teal-50 p-4 rounded-lg"
          >
            <h4 class="text-lg font-medium text-teal-900 mb-4">
              ✅ Eligibility Criteria
            </h4>
            <div>
              <label
                for="eligibility_criteria"
                class="block text-sm font-medium text-gray-700 mb-1"
                >Eligibility Criteria</label
              >
              <textarea
                id="eligibility_criteria"
                v-model="form.eligibility_criteria"
                rows="4"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="Describe who is eligible for services, age requirements, residency requirements, etc..."
              ></textarea>
            </div>
          </div>

          <!-- Submitter Information -->
          <div class="bg-gray-50 p-4 rounded-lg">
            <h4 class="text-lg font-medium text-gray-900 mb-4">
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
                  >Relationship to Regional Center *</label
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
                  <option value="Client/Consumer">Client/Consumer</option>
                  <option value="Family Member">Family Member</option>
                  <option value="Service Provider">Service Provider</option>
                  <option value="Community Partner">Community Partner</option>
                  <option value="Healthcare Provider">
                    Healthcare Provider
                  </option>
                  <option value="Advocate">Advocate</option>
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
                <div class="font-medium">Update Submitted Successfully!</div>
                <div class="text-sm">
                  Thank you for helping us keep regional center information
                  accurate.
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
              <span v-else>📋</span>
              <span>{{
                loading
                  ? "Submitting..."
                  : submitted
                  ? "Submitted"
                  : "Submit Update"
              }}</span>
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, reactive, onMounted } from "vue";
import { apiService } from "@/services/api";

export default {
  name: "RegionalCenterUpdateForm",
  props: {
    isVisible: {
      type: Boolean,
      default: false,
    },
    preSelectedCenter: {
      type: Object,
      default: null,
    },
  },
  emits: ["close", "submitted"],
  setup(props, { emit }) {
    const form = reactive({
      regional_center_name: "",
      other_center_name: "",
      regional_center_id: "",
      update_type: "",
      update_description: "",

      // Contact info
      new_phone: "",
      new_email: "",
      new_website: "",
      new_address: "",

      // Services
      services_offered: "",
      special_programs: "",

      // Hours
      operating_hours: "",

      // Eligibility
      eligibility_criteria: "",

      // Submitter
      submitter_name: "",
      submitter_email: "",
      relationship: "",
    });

    const loading = ref(false);
    const error = ref("");
    const submitted = ref(false);
    const regionalCenters = ref([]);

    const loadRegionalCenters = async () => {
      try {
        const centers = await apiService.getRegionalCenters();
        regionalCenters.value = centers || [];
      } catch (err) {
        console.error("Error loading regional centers:", err);
      }
    };

    const onRegionalCenterChange = () => {
      const selectedCenter = regionalCenters.value.find(
        (center) => center.regional_center === form.regional_center_name
      );
      if (selectedCenter) {
        form.regional_center_id = selectedCenter.id.toString();
      }
    };

    const resetForm = () => {
      Object.assign(form, {
        regional_center_name: "",
        other_center_name: "",
        regional_center_id: "",
        update_type: "",
        update_description: "",
        new_phone: "",
        new_email: "",
        new_website: "",
        new_address: "",
        services_offered: "",
        special_programs: "",
        operating_hours: "",
        eligibility_criteria: "",
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
        // Enhanced validation
        const missingFields = [];
        if (!form.regional_center_name) missingFields.push("Regional Center");
        if (form.regional_center_name === "other" && !form.other_center_name)
          missingFields.push("Regional Center Name");
        if (!form.update_type) missingFields.push("Update Type");
        if (!form.update_description.trim())
          missingFields.push("Update Description");
        if (!form.submitter_name.trim()) missingFields.push("Your Name");
        if (!form.submitter_email.trim()) missingFields.push("Your Email");
        if (!form.relationship) missingFields.push("Relationship");

        if (missingFields.length > 0) {
          error.value = `Please fill out the following required fields: ${missingFields.join(
            ", "
          )}`;
          loading.value = false;
          return;
        }

        const submissionData = {
          ...form,
          regional_center_name:
            form.regional_center_name === "other"
              ? form.other_center_name
              : form.regional_center_name,
          submitted_at: new Date().toISOString(),
        };

        await apiService.submitRegionalCenterUpdate(submissionData);
        submitted.value = true;
        emit("submitted", submissionData);

        setTimeout(() => {
          if (submitted.value) {
            closeModal();
          }
        }, 3000);
      } catch (err) {
        console.error("Error submitting regional center update:", err);
        error.value =
          err.message || "Failed to submit update. Please try again.";
      } finally {
        loading.value = false;
      }
    };

    onMounted(() => {
      loadRegionalCenters();

      // Pre-populate if center is provided
      if (props.preSelectedCenter) {
        form.regional_center_name = props.preSelectedCenter.regional_center;
        form.regional_center_id = props.preSelectedCenter.id.toString();
      }
    });

    return {
      form,
      loading,
      error,
      submitted,
      regionalCenters,
      onRegionalCenterChange,
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

.overflow-y-auto {
  scrollbar-width: thin;
  scrollbar-color: rgba(156, 163, 175, 0.5) transparent;
}

.overflow-y-auto::-webkit-scrollbar {
  width: 6px;
}

.overflow-y-auto::-webkit-scrollbar-track {
  background: transparent;
}

.overflow-y-auto::-webkit-scrollbar-thumb {
  background-color: rgba(156, 163, 175, 0.5);
  border-radius: 3px;
}
</style>
