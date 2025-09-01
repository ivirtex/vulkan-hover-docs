# VK\_FUCHSIA\_external\_memory(3) Manual Page

## Name

VK\_FUCHSIA\_external\_memory - device extension



## [](#_registered_extension_number)Registered Extension Number

365

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

     [VK\_KHR\_external\_memory\_capabilities](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_memory_capabilities.html)  
     and  
     [VK\_KHR\_external\_memory](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_memory.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_contact)Contact

- John Rosasco [\[GitHub\]rosasco](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_FUCHSIA_external_memory%5D%20%40rosasco%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_FUCHSIA_external_memory%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2021-03-01

**IP Status**

No known IP claims.

**Contributors**

- Craig Stout, Google
- John Bauman, Google
- John Rosasco, Google

## [](#_description)Description

Vulkan apps may wish to export or import device memory handles to or from other logical devices, instances or APIs.

This memory sharing can eliminate copies of memory buffers when different subsystems need to interoperate on them. Sharing memory buffers may also facilitate a better distribution of processing workload for more complex memory manipulation pipelines.

## [](#_new_commands)New Commands

- [vkGetMemoryZirconHandleFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetMemoryZirconHandleFUCHSIA.html)
- [vkGetMemoryZirconHandlePropertiesFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetMemoryZirconHandlePropertiesFUCHSIA.html)

## [](#_new_structures)New Structures

- [VkMemoryGetZirconHandleInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryGetZirconHandleInfoFUCHSIA.html)
- [VkMemoryZirconHandlePropertiesFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryZirconHandlePropertiesFUCHSIA.html)
- Extending [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateInfo.html):
  
  - [VkImportMemoryZirconHandleInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportMemoryZirconHandleInfoFUCHSIA.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_FUCHSIA_EXTERNAL_MEMORY_EXTENSION_NAME`
- `VK_FUCHSIA_EXTERNAL_MEMORY_SPEC_VERSION`
- Extending [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html):
  
  - `VK_EXTERNAL_MEMORY_HANDLE_TYPE_ZIRCON_VMO_BIT_FUCHSIA`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_IMPORT_MEMORY_ZIRCON_HANDLE_INFO_FUCHSIA`
  - `VK_STRUCTURE_TYPE_MEMORY_GET_ZIRCON_HANDLE_INFO_FUCHSIA`
  - `VK_STRUCTURE_TYPE_MEMORY_ZIRCON_HANDLE_PROPERTIES_FUCHSIA`

## [](#_issues)Issues

See `VK_KHR_external_memory` issues list for further information.

## [](#_version_history)Version History

- Revision 1, 2021-03-01 (John Rosasco)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_FUCHSIA_external_memory).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0