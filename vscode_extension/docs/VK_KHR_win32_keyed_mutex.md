# VK\_KHR\_win32\_keyed\_mutex(3) Manual Page

## Name

VK\_KHR\_win32\_keyed\_mutex - device extension



## [](#_registered_extension_number)Registered Extension Number

76

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_external\_memory\_win32](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_memory_win32.html)

## [](#_contact)Contact

- Carsten Rohde [\[GitHub\]crohde](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_win32_keyed_mutex%5D%20%40crohde%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_win32_keyed_mutex%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2016-10-21

**IP Status**

No known IP claims.

**Contributors**

- James Jones, NVIDIA
- Jeff Juliano, NVIDIA
- Carsten Rohde, NVIDIA

## [](#_description)Description

Applications that wish to import Direct3D 11 memory objects into the Vulkan API may wish to use the native keyed mutex mechanism to synchronize access to the memory between Vulkan and Direct3D. This extension provides a way for an application to access the keyed mutex associated with an imported Vulkan memory object when submitting command buffers to a queue.

## [](#_new_structures)New Structures

- Extending [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubmitInfo.html), [VkSubmitInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubmitInfo2.html):
  
  - [VkWin32KeyedMutexAcquireReleaseInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWin32KeyedMutexAcquireReleaseInfoKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_WIN32_KEYED_MUTEX_EXTENSION_NAME`
- `VK_KHR_WIN32_KEYED_MUTEX_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_WIN32_KEYED_MUTEX_ACQUIRE_RELEASE_INFO_KHR`

## [](#_version_history)Version History

- Revision 1, 2016-10-21 (James Jones)
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_win32_keyed_mutex).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0