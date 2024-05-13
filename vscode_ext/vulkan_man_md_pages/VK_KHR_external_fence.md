# VK_KHR_external_fence(3) Manual Page

## Name

VK_KHR_external_fence - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

114

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_external_fence_capabilities](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_fence_capabilities.html)  

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.1-promotions"
  target="_blank" rel="noopener">Vulkan 1.1</a>

## <a href="#_contact" class="anchor"></a>Contact

- Jesse Hall <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_external_fence%5D%20@critsec%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_external_fence%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>critsec</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2017-05-08

**IP Status**  
No known IP claims.

**Contributors**  
- Jesse Hall, Google

- James Jones, NVIDIA

- Jeff Juliano, NVIDIA

- Cass Everitt, Oculus

- Contributors to
  [`VK_KHR_external_semaphore`](VK_KHR_external_semaphore.html)

## <a href="#_description" class="anchor"></a>Description

An application using external memory may wish to synchronize access to
that memory using fences. This extension enables an application to
create fences from which non-Vulkan handles that reference the
underlying synchronization primitive can be exported.

## <a href="#_promotion_to_vulkan_1_1" class="anchor"></a>Promotion to Vulkan 1.1

All functionality in this extension is included in core Vulkan 1.1, with
the KHR suffix omitted. The original type, enum and command names are
still available as aliases of the core functionality.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkFenceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFenceCreateInfo.html):

  - [VkExportFenceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportFenceCreateInfoKHR.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkFenceImportFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFenceImportFlagBitsKHR.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkFenceImportFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFenceImportFlagsKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_EXTERNAL_FENCE_EXTENSION_NAME`

- `VK_KHR_EXTERNAL_FENCE_SPEC_VERSION`

- Extending [VkFenceImportFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFenceImportFlagBits.html):

  - `VK_FENCE_IMPORT_TEMPORARY_BIT_KHR`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_EXPORT_FENCE_CREATE_INFO_KHR`

## <a href="#_issues" class="anchor"></a>Issues

This extension borrows concepts, semantics, and language from
[`VK_KHR_external_semaphore`](VK_KHR_external_semaphore.html). That
extension’s issues apply equally to this extension.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2017-05-08 (Jesse Hall)

  - Initial revision

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_external_fence"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
