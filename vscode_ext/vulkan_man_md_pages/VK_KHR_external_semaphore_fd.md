# VK_KHR_external_semaphore_fd(3) Manual Page

## Name

VK_KHR_external_semaphore_fd - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

80

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_external_semaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_semaphore.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_contact" class="anchor"></a>Contact

- James Jones <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_external_semaphore_fd%5D%20@cubanismo%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_external_semaphore_fd%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>cubanismo</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2016-10-21

**IP Status**  
No known IP claims.

**Contributors**  
- Jesse Hall, Google

- James Jones, NVIDIA

- Jeff Juliano, NVIDIA

- Carsten Rohde, NVIDIA

## <a href="#_description" class="anchor"></a>Description

An application using external memory may wish to synchronize access to
that memory using semaphores. This extension enables an application to
export semaphore payload to and import semaphore payload from POSIX file
descriptors.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkGetSemaphoreFdKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetSemaphoreFdKHR.html)

- [vkImportSemaphoreFdKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkImportSemaphoreFdKHR.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkImportSemaphoreFdInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportSemaphoreFdInfoKHR.html)

- [VkSemaphoreGetFdInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreGetFdInfoKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_EXTERNAL_SEMAPHORE_FD_EXTENSION_NAME`

- `VK_KHR_EXTERNAL_SEMAPHORE_FD_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_IMPORT_SEMAPHORE_FD_INFO_KHR`

  - `VK_STRUCTURE_TYPE_SEMAPHORE_GET_FD_INFO_KHR`

## <a href="#_issues" class="anchor"></a>Issues

1\) Does the application need to close the file descriptor returned by
[vkGetSemaphoreFdKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetSemaphoreFdKHR.html)?

**RESOLVED**: Yes, unless it is passed back in to a driver instance to
import the semaphore. A successful get call transfers ownership of the
file descriptor to the application, and a successful import transfers it
back to the driver. Destroying the original semaphore object will not
close the file descriptor or remove its reference to the underlying
semaphore resource associated with it.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2016-10-21 (Jesse Hall)

  - Initial revision

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_external_semaphore_fd"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
