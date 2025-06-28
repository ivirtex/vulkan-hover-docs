# VK\_KHR\_external\_semaphore(3) Manual Page

## Name

VK\_KHR\_external\_semaphore - device extension



## [](#_registered_extension_number)Registered Extension Number

78

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_external\_semaphore\_capabilities](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_semaphore_capabilities.html)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.1](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.1-promotions)

## [](#_contact)Contact

- James Jones [\[GitHub\]cubanismo](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_external_semaphore%5D%20%40cubanismo%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_external_semaphore%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2016-10-21

**IP Status**

No known IP claims.

**Contributors**

- Faith Ekstrand, Intel
- Jesse Hall, Google
- Tobias Hector, Imagination Technologies
- James Jones, NVIDIA
- Jeff Juliano, NVIDIA
- Matthew Netsch, Qualcomm Technologies, Inc.
- Ray Smith, ARM
- Lina Versace, Google

## [](#_description)Description

An application using external memory may wish to synchronize access to that memory using semaphores. This extension enables an application to create semaphores from which non-Vulkan handles that reference the underlying synchronization primitive can be exported.

## [](#_promotion_to_vulkan_1_1)Promotion to Vulkan 1.1

All functionality in this extension is included in core Vulkan 1.1, with the KHR suffix omitted. The original type, enum, and command names are still available as aliases of the core functionality.

## [](#_new_structures)New Structures

- Extending [VkSemaphoreCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreCreateInfo.html):
  
  - [VkExportSemaphoreCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportSemaphoreCreateInfoKHR.html)

## [](#_new_enums)New Enums

- [VkSemaphoreImportFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreImportFlagBitsKHR.html)

## [](#_new_bitmasks)New Bitmasks

- [VkSemaphoreImportFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreImportFlagsKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_EXTERNAL_SEMAPHORE_EXTENSION_NAME`
- `VK_KHR_EXTERNAL_SEMAPHORE_SPEC_VERSION`
- Extending [VkSemaphoreImportFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreImportFlagBits.html):
  
  - `VK_SEMAPHORE_IMPORT_TEMPORARY_BIT_KHR`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_EXPORT_SEMAPHORE_CREATE_INFO_KHR`

## [](#_issues)Issues

1\) Should there be restrictions on what side effects can occur when waiting on imported semaphores that are in an invalid state?

**RESOLVED**: Yes. Normally, validating such state would be the responsibility of the application, and the implementation would be free to enter an undefined state if valid usage rules were violated. However, this could cause security concerns when using imported semaphores, as it would require the importing application to trust the exporting application to ensure the state is valid. Requiring this level of trust is undesirable for many potential use cases.

2\) Must implementations validate external handles the application provides as input to semaphore state import operations?

**RESOLVED**: Implementations must return an error to the application if the provided semaphore state handle cannot be used to complete the requested import operation. However, implementations need not validate handles are of the exact type specified by the application.

## [](#_version_history)Version History

- Revision 1, 2016-10-21 (James Jones)
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_external_semaphore)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0