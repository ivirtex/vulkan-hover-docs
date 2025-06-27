# VK_EXT_metal_objects(3) Manual Page

## Name

VK_EXT_metal_objects - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

312

## <a href="#_revision" class="anchor"></a>Revision

2

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_contact" class="anchor"></a>Contact

- Bill Hollings <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_metal_objects%5D%20@billhollings%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_metal_objects%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>billhollings</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_EXT_metal_objects](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_metal_objects.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2024-04-04

**IP Status**  
No known IP claims.

**Contributors**  
- Bill Hollings, The Brenwill Workshop Ltd.

- Dzmitry Malyshau, Mozilla Corp.

## <a href="#_description" class="anchor"></a>Description

In a Vulkan implementation that is layered on top of Metal on Apple
device platforms, this extension provides the ability to import and
export the underlying Metal objects associated with specific Vulkan
objects.

As detailed in the [extension proposal
document](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_metal_objects.adoc),
this extension adds one new Vulkan command,
[vkExportMetalObjectsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkExportMetalObjectsEXT.html), to export
underlying Metal objects from Vulkan objects, and supports importing the
appropriate existing Metal objects when creating Vulkan objects of types
[VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html), [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html),
[VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html), and [VkEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkEvent.html),

The intent is that this extension will be advertised and supported only
on implementations that are layered on top of Metal on Apple device
platforms.

## <a href="#_new_base_types" class="anchor"></a>New Base Types

- [IOSurfaceRef](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/IOSurfaceRef.html)

- [MTLBuffer_id](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/MTLBuffer_id.html)

- [MTLCommandQueue_id](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/MTLCommandQueue_id.html)

- [MTLDevice_id](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/MTLDevice_id.html)

- [MTLSharedEvent_id](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/MTLSharedEvent_id.html)

- [MTLTexture_id](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/MTLTexture_id.html)

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkExportMetalObjectsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkExportMetalObjectsEXT.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkExportMetalObjectsInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalObjectsInfoEXT.html)

- Extending
  [VkExportMetalObjectsInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalObjectsInfoEXT.html):

  - [VkExportMetalBufferInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalBufferInfoEXT.html)

  - [VkExportMetalCommandQueueInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalCommandQueueInfoEXT.html)

  - [VkExportMetalDeviceInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalDeviceInfoEXT.html)

  - [VkExportMetalIOSurfaceInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalIOSurfaceInfoEXT.html)

  - [VkExportMetalSharedEventInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalSharedEventInfoEXT.html)

  - [VkExportMetalTextureInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalTextureInfoEXT.html)

- Extending [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html):

  - [VkImportMetalIOSurfaceInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportMetalIOSurfaceInfoEXT.html)

  - [VkImportMetalTextureInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportMetalTextureInfoEXT.html)

- Extending [VkInstanceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstanceCreateInfo.html),
  [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateInfo.html),
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html),
  [VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewCreateInfo.html),
  [VkBufferViewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferViewCreateInfo.html),
  [VkSemaphoreCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreCreateInfo.html),
  [VkEventCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkEventCreateInfo.html):

  - [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalObjectCreateInfoEXT.html)

- Extending [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateInfo.html):

  - [VkImportMetalBufferInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportMetalBufferInfoEXT.html)

- Extending [VkSemaphoreCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreCreateInfo.html),
  [VkEventCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkEventCreateInfo.html):

  - [VkImportMetalSharedEventInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportMetalSharedEventInfoEXT.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkExportMetalObjectTypeFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalObjectTypeFlagBitsEXT.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkExportMetalObjectTypeFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalObjectTypeFlagsEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_METAL_OBJECTS_EXTENSION_NAME`

- `VK_EXT_METAL_OBJECTS_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

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

## <a href="#_issues" class="anchor"></a>Issues

None.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2022-05-28 (Bill Hollings)

  - Initial draft.

  - Incorporated feedback from review by the Vulkan Working Group.
    Renamed many structures, moved import/export of MTLBuffer to
    VkDeviceMemory, added export of MTLSharedEvent, added import of
    MTLSharedEvent for VkSemaphore and VkEvent, and changed used bit
    mask fields to individual bit fields to simplify Valid Usage rules.

- Revision 2, 2024-04-04 (Bill Hollings)

  - Add an `__unsafe_unretained` ownership qualifier to all Metal object
    declarations, to support Automatic Reference Counting (ARC) on Apple
    devices.

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_metal_objects"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
