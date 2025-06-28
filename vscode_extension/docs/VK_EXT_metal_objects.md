# VK\_EXT\_metal\_objects(3) Manual Page

## Name

VK\_EXT\_metal\_objects - device extension



## [](#_registered_extension_number)Registered Extension Number

312

## [](#_revision)Revision

2

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_contact)Contact

- Bill Hollings [\[GitHub\]billhollings](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_metal_objects%5D%20%40billhollings%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_metal_objects%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_EXT\_metal\_objects](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_metal_objects.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2024-04-04

**IP Status**

No known IP claims.

**Contributors**

- Bill Hollings, The Brenwill Workshop Ltd.
- Dzmitry Malyshau, Mozilla Corp.

## [](#_description)Description

In a Vulkan implementation that is layered on top of Metal on Apple device platforms, this extension provides the ability to import and export the underlying Metal objects associated with specific Vulkan objects.

As detailed in the [extension proposal document](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_metal_objects.adoc), this extension adds one new Vulkan command, [vkExportMetalObjectsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkExportMetalObjectsEXT.html), to export underlying Metal objects from Vulkan objects, and supports importing the appropriate existing Metal objects when creating Vulkan objects of types [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html), [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html), [VkSemaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphore.html), and [VkEvent](https://registry.khronos.org/vulkan/specs/latest/man/html/VkEvent.html),

The intent is that this extension will be advertised and supported only on implementations that are layered on top of Metal on Apple device platforms.

## [](#_new_base_types)New Base Types

- [IOSurfaceRef](https://registry.khronos.org/vulkan/specs/latest/man/html/IOSurfaceRef.html)
- [MTLBuffer\_id](https://registry.khronos.org/vulkan/specs/latest/man/html/MTLBuffer_id.html)
- [MTLCommandQueue\_id](https://registry.khronos.org/vulkan/specs/latest/man/html/MTLCommandQueue_id.html)
- [MTLDevice\_id](https://registry.khronos.org/vulkan/specs/latest/man/html/MTLDevice_id.html)
- [MTLSharedEvent\_id](https://registry.khronos.org/vulkan/specs/latest/man/html/MTLSharedEvent_id.html)
- [MTLTexture\_id](https://registry.khronos.org/vulkan/specs/latest/man/html/MTLTexture_id.html)

## [](#_new_commands)New Commands

- [vkExportMetalObjectsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkExportMetalObjectsEXT.html)

## [](#_new_structures)New Structures

- [VkExportMetalObjectsInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalObjectsInfoEXT.html)
- Extending [VkExportMetalObjectsInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalObjectsInfoEXT.html):
  
  - [VkExportMetalBufferInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalBufferInfoEXT.html)
  - [VkExportMetalCommandQueueInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalCommandQueueInfoEXT.html)
  - [VkExportMetalDeviceInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalDeviceInfoEXT.html)
  - [VkExportMetalIOSurfaceInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalIOSurfaceInfoEXT.html)
  - [VkExportMetalSharedEventInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalSharedEventInfoEXT.html)
  - [VkExportMetalTextureInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalTextureInfoEXT.html)
- Extending [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html):
  
  - [VkImportMetalIOSurfaceInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportMetalIOSurfaceInfoEXT.html)
  - [VkImportMetalTextureInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportMetalTextureInfoEXT.html)
- Extending [VkInstanceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstanceCreateInfo.html), [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateInfo.html), [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html), [VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewCreateInfo.html), [VkBufferViewCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferViewCreateInfo.html), [VkSemaphoreCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreCreateInfo.html), [VkEventCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkEventCreateInfo.html):
  
  - [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalObjectCreateInfoEXT.html)
- Extending [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateInfo.html):
  
  - [VkImportMetalBufferInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportMetalBufferInfoEXT.html)
- Extending [VkSemaphoreCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreCreateInfo.html), [VkEventCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkEventCreateInfo.html):
  
  - [VkImportMetalSharedEventInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportMetalSharedEventInfoEXT.html)

## [](#_new_enums)New Enums

- [VkExportMetalObjectTypeFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalObjectTypeFlagBitsEXT.html)

## [](#_new_bitmasks)New Bitmasks

- [VkExportMetalObjectTypeFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalObjectTypeFlagsEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_METAL_OBJECTS_EXTENSION_NAME`
- `VK_EXT_METAL_OBJECTS_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_EXPORT_METAL_BUFFER_INFO_EXT`
  - `VK_STRUCTURE_TYPE_EXPORT_METAL_COMMAND_QUEUE_INFO_EXT`
  - `VK_STRUCTURE_TYPE_EXPORT_METAL_DEVICE_INFO_EXT`
  - `VK_STRUCTURE_TYPE_EXPORT_METAL_IO_SURFACE_INFO_EXT`
  - `VK_STRUCTURE_TYPE_EXPORT_METAL_OBJECTS_INFO_EXT`
  - `VK_STRUCTURE_TYPE_EXPORT_METAL_OBJECT_CREATE_INFO_EXT`
  - `VK_STRUCTURE_TYPE_EXPORT_METAL_SHARED_EVENT_INFO_EXT`
  - `VK_STRUCTURE_TYPE_EXPORT_METAL_TEXTURE_INFO_EXT`
  - `VK_STRUCTURE_TYPE_IMPORT_METAL_BUFFER_INFO_EXT`
  - `VK_STRUCTURE_TYPE_IMPORT_METAL_IO_SURFACE_INFO_EXT`
  - `VK_STRUCTURE_TYPE_IMPORT_METAL_SHARED_EVENT_INFO_EXT`
  - `VK_STRUCTURE_TYPE_IMPORT_METAL_TEXTURE_INFO_EXT`

## [](#_issues)Issues

None.

## [](#_version_history)Version History

- Revision 1, 2022-05-28 (Bill Hollings)
  
  - Initial draft.
  - Incorporated feedback from review by the Vulkan Working Group. Renamed many structures, moved import/export of MTLBuffer to VkDeviceMemory, added export of MTLSharedEvent, added import of MTLSharedEvent for VkSemaphore and VkEvent, and changed used bit mask fields to individual bit fields to simplify Valid Usage rules.
- Revision 2, 2024-04-04 (Bill Hollings)
  
  - Add an `__unsafe_unretained` ownership qualifier to all Metal object declarations, to support Automatic Reference Counting (ARC) on Apple devices.

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_metal_objects)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0