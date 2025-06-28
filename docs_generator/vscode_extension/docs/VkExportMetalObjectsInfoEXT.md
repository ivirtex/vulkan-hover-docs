# VkExportMetalObjectsInfoEXT(3) Manual Page

## Name

VkExportMetalObjectsInfoEXT - Structure whose pNext chain identifies Vulkan objects and corresponding Metal objects



## [](#_c_specification)C Specification

The `VkExportMetalObjectsInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_metal_objects
typedef struct VkExportMetalObjectsInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
} VkExportMetalObjectsInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.

## [](#_description)Description

Valid Usage

- [](#VUID-VkExportMetalObjectsInfoEXT-pNext-06791)VUID-VkExportMetalObjectsInfoEXT-pNext-06791  
  If the `pNext` chain includes a [VkExportMetalDeviceInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalDeviceInfoEXT.html) structure, the [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html) **must** have been created with `VK_EXPORT_METAL_OBJECT_TYPE_METAL_DEVICE_BIT_EXT` in the `exportObjectType` member of a [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalObjectCreateInfoEXT.html) structure in the `pNext` chain of the [VkInstanceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstanceCreateInfo.html) structure in the [vkCreateInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateInstance.html) command
- [](#VUID-VkExportMetalObjectsInfoEXT-pNext-06792)VUID-VkExportMetalObjectsInfoEXT-pNext-06792  
  If the `pNext` chain includes a [VkExportMetalCommandQueueInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalCommandQueueInfoEXT.html) structure, the [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html) **must** have been created with `VK_EXPORT_METAL_OBJECT_TYPE_METAL_COMMAND_QUEUE_BIT_EXT` in the `exportObjectType` member of a [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalObjectCreateInfoEXT.html) structure in the `pNext` chain of the [VkInstanceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstanceCreateInfo.html) structure in the [vkCreateInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateInstance.html) command
- [](#VUID-VkExportMetalObjectsInfoEXT-pNext-06793)VUID-VkExportMetalObjectsInfoEXT-pNext-06793  
  If the `pNext` chain includes a [VkExportMetalBufferInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalBufferInfoEXT.html) structure, the [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) in its `memory` member **must** have been allocated with `VK_EXPORT_METAL_OBJECT_TYPE_METAL_BUFFER_BIT_EXT` in the `exportObjectType` member of a [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalObjectCreateInfoEXT.html) structure in the `pNext` chain of the [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateInfo.html) structure in the [vkAllocateMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/vkAllocateMemory.html) command
- [](#VUID-VkExportMetalObjectsInfoEXT-pNext-06794)VUID-VkExportMetalObjectsInfoEXT-pNext-06794  
  If the `pNext` chain includes a [VkExportMetalTextureInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalTextureInfoEXT.html) structure, exactly one of its `image`, `imageView`, or `bufferView` members **must** not be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-VkExportMetalObjectsInfoEXT-pNext-06795)VUID-VkExportMetalObjectsInfoEXT-pNext-06795  
  If the `pNext` chain includes a [VkExportMetalTextureInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalTextureInfoEXT.html) structure, and its `image` member is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), the [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) in its `image` member **must** have been created with `VK_EXPORT_METAL_OBJECT_TYPE_METAL_TEXTURE_BIT_EXT` in the `exportObjectType` member of a [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalObjectCreateInfoEXT.html) structure in the `pNext` chain of the [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) structure in the [vkCreateImage](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateImage.html) command
- [](#VUID-VkExportMetalObjectsInfoEXT-pNext-06796)VUID-VkExportMetalObjectsInfoEXT-pNext-06796  
  If the `pNext` chain includes a [VkExportMetalTextureInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalTextureInfoEXT.html) structure, and its `imageView` member is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), the [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) in its `imageView` member **must** have been created with `VK_EXPORT_METAL_OBJECT_TYPE_METAL_TEXTURE_BIT_EXT` in the `exportObjectType` member of a [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalObjectCreateInfoEXT.html) structure in the `pNext` chain of the [VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewCreateInfo.html) structure in the [vkCreateImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateImageView.html) command
- [](#VUID-VkExportMetalObjectsInfoEXT-pNext-06797)VUID-VkExportMetalObjectsInfoEXT-pNext-06797  
  If the `pNext` chain includes a [VkExportMetalTextureInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalTextureInfoEXT.html) structure, and its `bufferView` member is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), the [VkBufferView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferView.html) in its `bufferView` member **must** have been created with `VK_EXPORT_METAL_OBJECT_TYPE_METAL_TEXTURE_BIT_EXT` in the `exportObjectType` member of a [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalObjectCreateInfoEXT.html) structure in the `pNext` chain of the [VkBufferViewCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferViewCreateInfo.html) structure in the [vkCreateBufferView](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateBufferView.html) command
- [](#VUID-VkExportMetalObjectsInfoEXT-pNext-06798)VUID-VkExportMetalObjectsInfoEXT-pNext-06798  
  If the `pNext` chain includes a [VkExportMetalTextureInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalTextureInfoEXT.html) structure, and if either its `image` or `imageView` member is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), then `plane` **must** be `VK_IMAGE_ASPECT_PLANE_0_BIT`, `VK_IMAGE_ASPECT_PLANE_1_BIT`, or `VK_IMAGE_ASPECT_PLANE_2_BIT`
- [](#VUID-VkExportMetalObjectsInfoEXT-pNext-06799)VUID-VkExportMetalObjectsInfoEXT-pNext-06799  
  If the `pNext` chain includes a [VkExportMetalTextureInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalTextureInfoEXT.html) structure, and if the [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) in its `image` member does not have a [multi-planar format](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-multiplanar), then its `plane` member **must** be `VK_IMAGE_ASPECT_PLANE_0_BIT`
- [](#VUID-VkExportMetalObjectsInfoEXT-pNext-06800)VUID-VkExportMetalObjectsInfoEXT-pNext-06800  
  If the `pNext` chain includes a [VkExportMetalTextureInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalTextureInfoEXT.html) structure, and if the [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) in its `image` member has a [multi-planar format](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-multiplanar) with only two planes, then its `plane` member **must** not be `VK_IMAGE_ASPECT_PLANE_2_BIT`
- [](#VUID-VkExportMetalObjectsInfoEXT-pNext-06801)VUID-VkExportMetalObjectsInfoEXT-pNext-06801  
  If the `pNext` chain includes a [VkExportMetalTextureInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalTextureInfoEXT.html) structure, and if the [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) in its `imageView` member does not have a [multi-planar format](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-multiplanar), then its `plane` member **must** be `VK_IMAGE_ASPECT_PLANE_0_BIT`
- [](#VUID-VkExportMetalObjectsInfoEXT-pNext-06802)VUID-VkExportMetalObjectsInfoEXT-pNext-06802  
  If the `pNext` chain includes a [VkExportMetalTextureInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalTextureInfoEXT.html) structure, and if the [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) in its `imageView` member has a [multi-planar format](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-multiplanar) with only two planes, then its `plane` member **must** not be `VK_IMAGE_ASPECT_PLANE_2_BIT`
- [](#VUID-VkExportMetalObjectsInfoEXT-pNext-06803)VUID-VkExportMetalObjectsInfoEXT-pNext-06803  
  If the `pNext` chain includes a [VkExportMetalIOSurfaceInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalIOSurfaceInfoEXT.html) structure, the [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) in its `image` member **must** have been created with `VK_EXPORT_METAL_OBJECT_TYPE_METAL_IOSURFACE_BIT_EXT` in the `exportObjectType` member of a [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalObjectCreateInfoEXT.html) structure in the `pNext` chain of the [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) structure in the [vkCreateImage](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateImage.html) command
- [](#VUID-VkExportMetalObjectsInfoEXT-pNext-06804)VUID-VkExportMetalObjectsInfoEXT-pNext-06804  
  If the `pNext` chain includes a [VkExportMetalSharedEventInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalSharedEventInfoEXT.html) structure, exactly one of its `semaphore` or `event` members **must** not be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-VkExportMetalObjectsInfoEXT-pNext-06805)VUID-VkExportMetalObjectsInfoEXT-pNext-06805  
  If the `pNext` chain includes a [VkExportMetalSharedEventInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalSharedEventInfoEXT.html) structure, and its `semaphore` member is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), the [VkSemaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphore.html) in its `semaphore` member **must** have been created with `VK_EXPORT_METAL_OBJECT_TYPE_METAL_SHARED_EVENT_BIT_EXT` in the `exportObjectType` member of a [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalObjectCreateInfoEXT.html) structure in the `pNext` chain of the [VkSemaphoreCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreCreateInfo.html) structure in the [vkCreateSemaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateSemaphore.html) command
- [](#VUID-VkExportMetalObjectsInfoEXT-pNext-06806)VUID-VkExportMetalObjectsInfoEXT-pNext-06806  
  If the `pNext` chain includes a [VkExportMetalSharedEventInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalSharedEventInfoEXT.html) structure, and its `event` member is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), the [VkEvent](https://registry.khronos.org/vulkan/specs/latest/man/html/VkEvent.html) in its `event` member **must** have been created with `VK_EXPORT_METAL_OBJECT_TYPE_METAL_SHARED_EVENT_BIT_EXT` in the `exportObjectType` member of a [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalObjectCreateInfoEXT.html) structure in the `pNext` chain of the [VkEventCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkEventCreateInfo.html) structure in the [vkCreateEvent](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateEvent.html) command

Valid Usage (Implicit)

- [](#VUID-VkExportMetalObjectsInfoEXT-sType-sType)VUID-VkExportMetalObjectsInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_EXPORT_METAL_OBJECTS_INFO_EXT`
- [](#VUID-VkExportMetalObjectsInfoEXT-pNext-pNext)VUID-VkExportMetalObjectsInfoEXT-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkExportMetalBufferInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalBufferInfoEXT.html), [VkExportMetalCommandQueueInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalCommandQueueInfoEXT.html), [VkExportMetalDeviceInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalDeviceInfoEXT.html), [VkExportMetalIOSurfaceInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalIOSurfaceInfoEXT.html), [VkExportMetalSharedEventInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalSharedEventInfoEXT.html), or [VkExportMetalTextureInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalTextureInfoEXT.html)
- [](#VUID-VkExportMetalObjectsInfoEXT-sType-unique)VUID-VkExportMetalObjectsInfoEXT-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique, with the exception of structures of type [VkExportMetalBufferInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalBufferInfoEXT.html), [VkExportMetalCommandQueueInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalCommandQueueInfoEXT.html), [VkExportMetalIOSurfaceInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalIOSurfaceInfoEXT.html), [VkExportMetalSharedEventInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalSharedEventInfoEXT.html), or [VkExportMetalTextureInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalTextureInfoEXT.html)

## [](#_see_also)See Also

[VK\_EXT\_metal\_objects](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_metal_objects.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkExportMetalObjectsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkExportMetalObjectsEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkExportMetalObjectsInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0