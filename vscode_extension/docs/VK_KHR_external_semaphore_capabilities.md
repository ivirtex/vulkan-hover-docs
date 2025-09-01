# VK\_KHR\_external\_semaphore\_capabilities(3) Manual Page

## Name

VK\_KHR\_external\_semaphore\_capabilities - instance extension



## [](#_registered_extension_number)Registered Extension Number

77

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.1](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.1-promotions)

## [](#_contact)Contact

- James Jones [\[GitHub\]cubanismo](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_external_semaphore_capabilities%5D%20%40cubanismo%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_external_semaphore_capabilities%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2016-10-20

**IP Status**

No known IP claims.

**Contributors**

- Jesse Hall, Google
- James Jones, NVIDIA
- Jeff Juliano, NVIDIA

## [](#_description)Description

An application may wish to reference device semaphores in multiple Vulkan logical devices or instances, in multiple processes, and/or in multiple APIs. This extension provides a set of capability queries and handle definitions that allow an application to determine what types of “external” semaphore handles an implementation supports for a given set of use cases.

## [](#_promotion_to_vulkan_1_1)Promotion to Vulkan 1.1

All functionality in this extension is included in core Vulkan 1.1, with the KHR suffix omitted. The original type, enum, and command names are still available as aliases of the core functionality.

## [](#_new_commands)New Commands

- [vkGetPhysicalDeviceExternalSemaphorePropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceExternalSemaphorePropertiesKHR.html)

## [](#_new_structures)New Structures

- [VkExternalSemaphorePropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalSemaphorePropertiesKHR.html)
- [VkPhysicalDeviceExternalSemaphoreInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceExternalSemaphoreInfoKHR.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceIDPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceIDPropertiesKHR.html)

## [](#_new_enums)New Enums

- [VkExternalSemaphoreFeatureFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalSemaphoreFeatureFlagBitsKHR.html)
- [VkExternalSemaphoreHandleTypeFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalSemaphoreHandleTypeFlagBitsKHR.html)

## [](#_new_bitmasks)New Bitmasks

- [VkExternalSemaphoreFeatureFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalSemaphoreFeatureFlagsKHR.html)
- [VkExternalSemaphoreHandleTypeFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalSemaphoreHandleTypeFlagsKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_EXTERNAL_SEMAPHORE_CAPABILITIES_EXTENSION_NAME`
- `VK_KHR_EXTERNAL_SEMAPHORE_CAPABILITIES_SPEC_VERSION`
- `VK_LUID_SIZE_KHR`
- Extending [VkExternalSemaphoreFeatureFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalSemaphoreFeatureFlagBits.html):
  
  - `VK_EXTERNAL_SEMAPHORE_FEATURE_EXPORTABLE_BIT_KHR`
  - `VK_EXTERNAL_SEMAPHORE_FEATURE_IMPORTABLE_BIT_KHR`
- Extending [VkExternalSemaphoreHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalSemaphoreHandleTypeFlagBits.html):
  
  - `VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_D3D12_FENCE_BIT_KHR`
  - `VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_FD_BIT_KHR`
  - `VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_BIT_KHR`
  - `VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT_KHR`
  - `VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_SYNC_FD_BIT_KHR`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_EXTERNAL_SEMAPHORE_PROPERTIES_KHR`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_SEMAPHORE_INFO_KHR`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ID_PROPERTIES_KHR`

## [](#_version_history)Version History

- Revision 1, 2016-10-20 (James Jones)
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_external_semaphore_capabilities).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0