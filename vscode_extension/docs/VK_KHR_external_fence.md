# VK\_KHR\_external\_fence(3) Manual Page

## Name

VK\_KHR\_external\_fence - device extension



## [](#_registered_extension_number)Registered Extension Number

114

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_external\_fence\_capabilities](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_fence_capabilities.html)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.1](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.1-promotions)

## [](#_contact)Contact

- Jesse Hall [\[GitHub\]critsec](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_external_fence%5D%20%40critsec%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_external_fence%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2017-05-08

**IP Status**

No known IP claims.

**Contributors**

- Jesse Hall, Google
- James Jones, NVIDIA
- Jeff Juliano, NVIDIA
- Cass Everitt, Oculus
- Contributors to `VK_KHR_external_semaphore`

## [](#_description)Description

An application using external memory may wish to synchronize access to that memory using fences. This extension enables an application to create fences from which non-Vulkan handles that reference the underlying synchronization primitive can be exported.

## [](#_promotion_to_vulkan_1_1)Promotion to Vulkan 1.1

All functionality in this extension is included in core Vulkan 1.1, with the KHR suffix omitted. The original type, enum, and command names are still available as aliases of the core functionality.

## [](#_new_structures)New Structures

- Extending [VkFenceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFenceCreateInfo.html):
  
  - [VkExportFenceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportFenceCreateInfoKHR.html)

## [](#_new_enums)New Enums

- [VkFenceImportFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFenceImportFlagBitsKHR.html)

## [](#_new_bitmasks)New Bitmasks

- [VkFenceImportFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFenceImportFlagsKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_EXTERNAL_FENCE_EXTENSION_NAME`
- `VK_KHR_EXTERNAL_FENCE_SPEC_VERSION`
- Extending [VkFenceImportFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFenceImportFlagBits.html):
  
  - `VK_FENCE_IMPORT_TEMPORARY_BIT_KHR`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_EXPORT_FENCE_CREATE_INFO_KHR`

## [](#_issues)Issues

This extension borrows concepts, semantics, and language from `VK_KHR_external_semaphore`. That extension’s issues apply equally to this extension.

## [](#_version_history)Version History

- Revision 1, 2017-05-08 (Jesse Hall)
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_external_fence)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0