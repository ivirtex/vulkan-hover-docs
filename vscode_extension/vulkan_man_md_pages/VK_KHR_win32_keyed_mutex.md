# VK_KHR_win32_keyed_mutex(3) Manual Page

## Name

VK_KHR_win32_keyed_mutex - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

76

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_external_memory_win32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_memory_win32.html)  

## <a href="#_contact" class="anchor"></a>Contact

- Carsten Rohde <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_win32_keyed_mutex%5D%20@crohde%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_win32_keyed_mutex%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>crohde</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2016-10-21

**IP Status**  
No known IP claims.

**Contributors**  
- James Jones, NVIDIA

- Jeff Juliano, NVIDIA

- Carsten Rohde, NVIDIA

## <a href="#_description" class="anchor"></a>Description

Applications that wish to import Direct3D 11 memory objects into the
Vulkan API may wish to use the native keyed mutex mechanism to
synchronize access to the memory between Vulkan and Direct3D. This
extension provides a way for an application to access the keyed mutex
associated with an imported Vulkan memory object when submitting command
buffers to a queue.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubmitInfo.html),
  [VkSubmitInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubmitInfo2.html):

  - [VkWin32KeyedMutexAcquireReleaseInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWin32KeyedMutexAcquireReleaseInfoKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_WIN32_KEYED_MUTEX_EXTENSION_NAME`

- `VK_KHR_WIN32_KEYED_MUTEX_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_WIN32_KEYED_MUTEX_ACQUIRE_RELEASE_INFO_KHR`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2016-10-21 (James Jones)

  - Initial revision

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_win32_keyed_mutex"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
