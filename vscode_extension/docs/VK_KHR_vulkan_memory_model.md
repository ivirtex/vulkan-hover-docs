# VK\_KHR\_vulkan\_memory\_model(3) Manual Page

## Name

VK\_KHR\_vulkan\_memory\_model - device extension



## [](#_registered_extension_number)Registered Extension Number

212

## [](#_revision)Revision

3

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_KHR\_vulkan\_memory\_model](https://github.khronos.org/SPIRV-Registry/extensions/KHR/SPV_KHR_vulkan_memory_model.html)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.2](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.2-promotions)

## [](#_contact)Contact

- Jeff Bolz [\[GitHub\]jeffbolznv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_vulkan_memory_model%5D%20%40jeffbolznv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_vulkan_memory_model%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2018-12-10

**IP Status**

No known IP claims.

**Contributors**

- Jeff Bolz, NVIDIA
- Alan Baker, Google
- Tobias Hector, AMD
- David Neto, Google
- Robert Simpson, Qualcomm Technologies, Inc.
- Brian Sumner, AMD

## [](#_description)Description

The [VK\_KHR\_vulkan\_memory\_model](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_vulkan_memory_model.html) extension allows use of the features guarded by the `VulkanMemoryModel`, `VulkanMemoryModelDeviceScope`, and `VulkanMemoryModelAvailabilityVisibilityChains` capabilities in shader modules. The [Vulkan Memory Model](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-model) formally defines how to synchronize memory accesses to the same memory locations performed by multiple shader invocations.

Note

Version 3 of the spec added a member (`vulkanMemoryModelAvailabilityVisibilityChains`) to [VkPhysicalDeviceVulkanMemoryModelFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVulkanMemoryModelFeaturesKHR.html), which is an incompatible change from version 2.

## [](#_promotion_to_vulkan_1_2)Promotion to Vulkan 1.2

All functionality in this extension is included in core Vulkan 1.2, with the KHR suffix omitted. However, if Vulkan 1.2 is supported and this extension is not, the `vulkanMemoryModel` capability is optional. The original type, enum, and command names are still available as aliases of the core functionality.

## [](#_promotion_to_vulkan_1_3)Promotion to Vulkan 1.3

If Vulkan 1.3 is supported, support for the `vulkanMemoryModel` and `vulkanMemoryModelDeviceScope` capabilities is required.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceVulkanMemoryModelFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVulkanMemoryModelFeaturesKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_VULKAN_MEMORY_MODEL_EXTENSION_NAME`
- `VK_KHR_VULKAN_MEMORY_MODEL_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_MEMORY_MODEL_FEATURES_KHR`

## [](#_new_spir_v_capabilities)New SPIR-V Capabilities

- [`VulkanMemoryModelKHR`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-VulkanMemoryModel)

## [](#_version_history)Version History

- Revision 1, 2018-06-24 (Jeff Bolz)
  
  - Initial draft
- Revision 3, 2018-12-10 (Jeff Bolz)
  
  - Add vulkanMemoryModelAvailabilityVisibilityChains member to the VkPhysicalDeviceVulkanMemoryModelFeaturesKHR structure.

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_vulkan_memory_model).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0