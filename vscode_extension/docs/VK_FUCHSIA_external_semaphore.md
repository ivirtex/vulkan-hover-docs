# VK\_FUCHSIA\_external\_semaphore(3) Manual Page

## Name

VK\_FUCHSIA\_external\_semaphore - device extension



## [](#_registered_extension_number)Registered Extension Number

366

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_external\_semaphore\_capabilities](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_semaphore_capabilities.html)  
and  
[VK\_KHR\_external\_semaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_semaphore.html)

## [](#_contact)Contact

- John Rosasco [\[GitHub\]rosasco](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_FUCHSIA_external_semaphore%5D%20%40rosasco%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_FUCHSIA_external_semaphore%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2021-03-08

**IP Status**

No known IP claims.

**Contributors**

- Craig Stout, Google
- John Bauman, Google
- John Rosasco, Google

## [](#_description)Description

An application using external memory may wish to synchronize access to that memory using semaphores. This extension enables an application to export semaphore payload to and import semaphore payload from Zircon event handles.

## [](#_new_commands)New Commands

- [vkGetSemaphoreZirconHandleFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetSemaphoreZirconHandleFUCHSIA.html)
- [vkImportSemaphoreZirconHandleFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/vkImportSemaphoreZirconHandleFUCHSIA.html)

## [](#_new_structures)New Structures

- [VkImportSemaphoreZirconHandleInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportSemaphoreZirconHandleInfoFUCHSIA.html)
- [VkSemaphoreGetZirconHandleInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreGetZirconHandleInfoFUCHSIA.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_FUCHSIA_EXTERNAL_SEMAPHORE_EXTENSION_NAME`
- `VK_FUCHSIA_EXTERNAL_SEMAPHORE_SPEC_VERSION`
- Extending [VkExternalSemaphoreHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalSemaphoreHandleTypeFlagBits.html):
  
  - `VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_ZIRCON_EVENT_BIT_FUCHSIA`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_IMPORT_SEMAPHORE_ZIRCON_HANDLE_INFO_FUCHSIA`
  - `VK_STRUCTURE_TYPE_SEMAPHORE_GET_ZIRCON_HANDLE_INFO_FUCHSIA`

## [](#_issues)Issues

1\) Does the application need to close the Zircon event handle returned by [vkGetSemaphoreZirconHandleFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetSemaphoreZirconHandleFUCHSIA.html)?

**RESOLVED**: Yes, unless it is passed back in to a driver instance to import the semaphore. A successful get call transfers ownership of the Zircon event handle to the application, and a successful import transfers it back to the driver. Destroying the original semaphore object will not close the Zircon event handle nor remove its reference to the underlying semaphore resource associated with it.

## [](#_version_history)Version History

- Revision 1, 2021-03-08 (John Rosasco)
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_FUCHSIA_external_semaphore).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0