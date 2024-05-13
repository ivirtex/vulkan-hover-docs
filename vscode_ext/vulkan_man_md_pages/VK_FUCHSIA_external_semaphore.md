# VK_FUCHSIA_external_semaphore(3) Manual Page

## Name

VK_FUCHSIA_external_semaphore - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

366

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_external_semaphore_capabilities](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_semaphore_capabilities.html)  
and  
[VK_KHR_external_semaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_semaphore.html)  

## <a href="#_contact" class="anchor"></a>Contact

- John Rosasco <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_FUCHSIA_external_semaphore%5D%20@rosasco%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_FUCHSIA_external_semaphore%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>rosasco</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2021-03-08

**IP Status**  
No known IP claims.

**Contributors**  
- Craig Stout, Google

- John Bauman, Google

- John Rosasco, Google

## <a href="#_description" class="anchor"></a>Description

An application using external memory may wish to synchronize access to
that memory using semaphores. This extension enables an application to
export semaphore payload to and import semaphore payload from Zircon
event handles.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkGetSemaphoreZirconHandleFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetSemaphoreZirconHandleFUCHSIA.html)

- [vkImportSemaphoreZirconHandleFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkImportSemaphoreZirconHandleFUCHSIA.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkImportSemaphoreZirconHandleInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportSemaphoreZirconHandleInfoFUCHSIA.html)

- [VkSemaphoreGetZirconHandleInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreGetZirconHandleInfoFUCHSIA.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_FUCHSIA_EXTERNAL_SEMAPHORE_EXTENSION_NAME`

- `VK_FUCHSIA_EXTERNAL_SEMAPHORE_SPEC_VERSION`

- Extending
  [VkExternalSemaphoreHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreHandleTypeFlagBits.html):

  - `VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_ZIRCON_EVENT_BIT_FUCHSIA`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_IMPORT_SEMAPHORE_ZIRCON_HANDLE_INFO_FUCHSIA`

  - `VK_STRUCTURE_TYPE_SEMAPHORE_GET_ZIRCON_HANDLE_INFO_FUCHSIA`

## <a href="#_issues" class="anchor"></a>Issues

1\) Does the application need to close the Zircon event handle returned
by
[vkGetSemaphoreZirconHandleFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetSemaphoreZirconHandleFUCHSIA.html)?

**RESOLVED**: Yes, unless it is passed back in to a driver instance to
import the semaphore. A successful get call transfers ownership of the
Zircon event handle to the application, and a successful import
transfers it back to the driver. Destroying the original semaphore
object will not close the Zircon event handle nor remove its reference
to the underlying semaphore resource associated with it.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2021-03-08 (John Rosasco)

  - Initial revision

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_FUCHSIA_external_semaphore"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
