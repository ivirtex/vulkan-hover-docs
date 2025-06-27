# VK_FUCHSIA_external_memory(3) Manual Page

## Name

VK_FUCHSIA_external_memory - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

365

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

    
[VK_KHR_external_memory_capabilities](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_memory_capabilities.html)  
     and  
     [VK_KHR_external_memory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_memory.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_contact" class="anchor"></a>Contact

- John Rosasco <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_FUCHSIA_external_memory%5D%20@rosasco%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_FUCHSIA_external_memory%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>rosasco</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2021-03-01

**IP Status**  
No known IP claims.

**Contributors**  
- Craig Stout, Google

- John Bauman, Google

- John Rosasco, Google

## <a href="#_description" class="anchor"></a>Description

Vulkan apps may wish to export or import device memory handles to or
from other logical devices, instances or APIs.

This memory sharing can eliminate copies of memory buffers when
different subsystems need to interoperate on them. Sharing memory
buffers may also facilitate a better distribution of processing workload
for more complex memory manipulation pipelines.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkGetMemoryZirconHandleFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetMemoryZirconHandleFUCHSIA.html)

- [vkGetMemoryZirconHandlePropertiesFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetMemoryZirconHandlePropertiesFUCHSIA.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkMemoryGetZirconHandleInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryGetZirconHandleInfoFUCHSIA.html)

- [VkMemoryZirconHandlePropertiesFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryZirconHandlePropertiesFUCHSIA.html)

- Extending [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateInfo.html):

  - [VkImportMemoryZirconHandleInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportMemoryZirconHandleInfoFUCHSIA.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_FUCHSIA_EXTERNAL_MEMORY_EXTENSION_NAME`

- `VK_FUCHSIA_EXTERNAL_MEMORY_SPEC_VERSION`

- Extending
  [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html):

  - `VK_EXTERNAL_MEMORY_HANDLE_TYPE_ZIRCON_VMO_BIT_FUCHSIA`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_IMPORT_MEMORY_ZIRCON_HANDLE_INFO_FUCHSIA`

  - `VK_STRUCTURE_TYPE_MEMORY_GET_ZIRCON_HANDLE_INFO_FUCHSIA`

  - `VK_STRUCTURE_TYPE_MEMORY_ZIRCON_HANDLE_PROPERTIES_FUCHSIA`

## <a href="#_issues" class="anchor"></a>Issues

See [`VK_KHR_external_memory`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_memory.html) issues list
for further information.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2021-03-01 (John Rosasco)

  - Initial draft

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_FUCHSIA_external_memory"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
