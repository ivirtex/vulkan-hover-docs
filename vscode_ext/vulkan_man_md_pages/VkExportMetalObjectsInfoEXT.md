# VkExportMetalObjectsInfoEXT(3) Manual Page

## Name

VkExportMetalObjectsInfoEXT - Structure whose pNext chain identifies
Vulkan objects and corresponding Metal objects



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkExportMetalObjectsInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_metal_objects
typedef struct VkExportMetalObjectsInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
} VkExportMetalObjectsInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkExportMetalObjectsInfoEXT-pNext-06791"
  id="VUID-VkExportMetalObjectsInfoEXT-pNext-06791"></a>
  VUID-VkExportMetalObjectsInfoEXT-pNext-06791  
  If the `pNext` chain includes a
  [VkExportMetalDeviceInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalDeviceInfoEXT.html)
  structure, the [VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html) **must** have been
  created with `VK_EXPORT_METAL_OBJECT_TYPE_METAL_DEVICE_BIT_EXT` in the
  `exportObjectType` member of a
  [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalObjectCreateInfoEXT.html)
  structure in the `pNext` chain of the
  [VkInstanceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstanceCreateInfo.html) structure in the
  [vkCreateInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateInstance.html) command

- <a href="#VUID-VkExportMetalObjectsInfoEXT-pNext-06792"
  id="VUID-VkExportMetalObjectsInfoEXT-pNext-06792"></a>
  VUID-VkExportMetalObjectsInfoEXT-pNext-06792  
  If the `pNext` chain includes a
  [VkExportMetalCommandQueueInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalCommandQueueInfoEXT.html)
  structure, the [VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html) **must** have been
  created with `VK_EXPORT_METAL_OBJECT_TYPE_METAL_COMMAND_QUEUE_BIT_EXT`
  in the `exportObjectType` member of a
  [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalObjectCreateInfoEXT.html)
  structure in the `pNext` chain of the
  [VkInstanceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstanceCreateInfo.html) structure in the
  [vkCreateInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateInstance.html) command

- <a href="#VUID-VkExportMetalObjectsInfoEXT-pNext-06793"
  id="VUID-VkExportMetalObjectsInfoEXT-pNext-06793"></a>
  VUID-VkExportMetalObjectsInfoEXT-pNext-06793  
  If the `pNext` chain includes a
  [VkExportMetalBufferInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalBufferInfoEXT.html)
  structure, the [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) in its `memory`
  member **must** have been allocated with
  `VK_EXPORT_METAL_OBJECT_TYPE_METAL_BUFFER_BIT_EXT` in the
  `exportObjectType` member of a
  [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalObjectCreateInfoEXT.html)
  structure in the `pNext` chain of the
  [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateInfo.html) structure in the
  [vkAllocateMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAllocateMemory.html) command

- <a href="#VUID-VkExportMetalObjectsInfoEXT-pNext-06794"
  id="VUID-VkExportMetalObjectsInfoEXT-pNext-06794"></a>
  VUID-VkExportMetalObjectsInfoEXT-pNext-06794  
  If the `pNext` chain includes a
  [VkExportMetalTextureInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalTextureInfoEXT.html)
  structure, exactly one of its `image`, `imageView`, or `bufferView`
  members **must** not be [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-VkExportMetalObjectsInfoEXT-pNext-06795"
  id="VUID-VkExportMetalObjectsInfoEXT-pNext-06795"></a>
  VUID-VkExportMetalObjectsInfoEXT-pNext-06795  
  If the `pNext` chain includes a
  [VkExportMetalTextureInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalTextureInfoEXT.html)
  structure, and its `image` member is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) in
  its `image` member **must** have been created with
  `VK_EXPORT_METAL_OBJECT_TYPE_METAL_TEXTURE_BIT_EXT` in the
  `exportObjectType` member of a
  [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalObjectCreateInfoEXT.html)
  structure in the `pNext` chain of the
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) structure in the
  [vkCreateImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateImage.html) command

- <a href="#VUID-VkExportMetalObjectsInfoEXT-pNext-06796"
  id="VUID-VkExportMetalObjectsInfoEXT-pNext-06796"></a>
  VUID-VkExportMetalObjectsInfoEXT-pNext-06796  
  If the `pNext` chain includes a
  [VkExportMetalTextureInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalTextureInfoEXT.html)
  structure, and its `imageView` member is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the
  [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) in its `imageView` member **must**
  have been created with
  `VK_EXPORT_METAL_OBJECT_TYPE_METAL_TEXTURE_BIT_EXT` in the
  `exportObjectType` member of a
  [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalObjectCreateInfoEXT.html)
  structure in the `pNext` chain of the
  [VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewCreateInfo.html) structure in the
  [vkCreateImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateImageView.html) command

- <a href="#VUID-VkExportMetalObjectsInfoEXT-pNext-06797"
  id="VUID-VkExportMetalObjectsInfoEXT-pNext-06797"></a>
  VUID-VkExportMetalObjectsInfoEXT-pNext-06797  
  If the `pNext` chain includes a
  [VkExportMetalTextureInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalTextureInfoEXT.html)
  structure, and its `bufferView` member is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the
  [VkBufferView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferView.html) in its `bufferView` member **must**
  have been created with
  `VK_EXPORT_METAL_OBJECT_TYPE_METAL_TEXTURE_BIT_EXT` in the
  `exportObjectType` member of a
  [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalObjectCreateInfoEXT.html)
  structure in the `pNext` chain of the
  [VkBufferViewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferViewCreateInfo.html) structure in the
  [vkCreateBufferView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateBufferView.html) command

- <a href="#VUID-VkExportMetalObjectsInfoEXT-pNext-06798"
  id="VUID-VkExportMetalObjectsInfoEXT-pNext-06798"></a>
  VUID-VkExportMetalObjectsInfoEXT-pNext-06798  
  If the `pNext` chain includes a
  [VkExportMetalTextureInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalTextureInfoEXT.html)
  structure, and if either its `image` or `imageView` member is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), then `plane` **must** be
  `VK_IMAGE_ASPECT_PLANE_0_BIT`, `VK_IMAGE_ASPECT_PLANE_1_BIT`, or
  `VK_IMAGE_ASPECT_PLANE_2_BIT`

- <a href="#VUID-VkExportMetalObjectsInfoEXT-pNext-06799"
  id="VUID-VkExportMetalObjectsInfoEXT-pNext-06799"></a>
  VUID-VkExportMetalObjectsInfoEXT-pNext-06799  
  If the `pNext` chain includes a
  [VkExportMetalTextureInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalTextureInfoEXT.html)
  structure, and if the [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) in its `image` member
  does not have a multi-planar format, then its `plane` member **must**
  be `VK_IMAGE_ASPECT_PLANE_0_BIT`

- <a href="#VUID-VkExportMetalObjectsInfoEXT-pNext-06800"
  id="VUID-VkExportMetalObjectsInfoEXT-pNext-06800"></a>
  VUID-VkExportMetalObjectsInfoEXT-pNext-06800  
  If the `pNext` chain includes a
  [VkExportMetalTextureInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalTextureInfoEXT.html)
  structure, and if the [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) in its `image` member
  has a multi-planar format with only two planes, then its `plane`
  member **must** not be `VK_IMAGE_ASPECT_PLANE_2_BIT`

- <a href="#VUID-VkExportMetalObjectsInfoEXT-pNext-06801"
  id="VUID-VkExportMetalObjectsInfoEXT-pNext-06801"></a>
  VUID-VkExportMetalObjectsInfoEXT-pNext-06801  
  If the `pNext` chain includes a
  [VkExportMetalTextureInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalTextureInfoEXT.html)
  structure, and if the [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) in its
  `imageView` member does not have a multi-planar format, then its
  `plane` member **must** be `VK_IMAGE_ASPECT_PLANE_0_BIT`

- <a href="#VUID-VkExportMetalObjectsInfoEXT-pNext-06802"
  id="VUID-VkExportMetalObjectsInfoEXT-pNext-06802"></a>
  VUID-VkExportMetalObjectsInfoEXT-pNext-06802  
  If the `pNext` chain includes a
  [VkExportMetalTextureInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalTextureInfoEXT.html)
  structure, and if the [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) in its
  `imageView` member has a multi-planar format with only two planes,
  then its `plane` member **must** not be `VK_IMAGE_ASPECT_PLANE_2_BIT`

- <a href="#VUID-VkExportMetalObjectsInfoEXT-pNext-06803"
  id="VUID-VkExportMetalObjectsInfoEXT-pNext-06803"></a>
  VUID-VkExportMetalObjectsInfoEXT-pNext-06803  
  If the `pNext` chain includes a
  [VkExportMetalIOSurfaceInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalIOSurfaceInfoEXT.html)
  structure, the [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) in its `image` member **must**
  have been created with
  `VK_EXPORT_METAL_OBJECT_TYPE_METAL_IOSURFACE_BIT_EXT` in the
  `exportObjectType` member of a
  [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalObjectCreateInfoEXT.html)
  structure in the `pNext` chain of the
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) structure in the
  [vkCreateImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateImage.html) command

- <a href="#VUID-VkExportMetalObjectsInfoEXT-pNext-06804"
  id="VUID-VkExportMetalObjectsInfoEXT-pNext-06804"></a>
  VUID-VkExportMetalObjectsInfoEXT-pNext-06804  
  If the `pNext` chain includes a
  [VkExportMetalSharedEventInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalSharedEventInfoEXT.html)
  structure, exactly one of its `semaphore` or `event` members **must**
  not be [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-VkExportMetalObjectsInfoEXT-pNext-06805"
  id="VUID-VkExportMetalObjectsInfoEXT-pNext-06805"></a>
  VUID-VkExportMetalObjectsInfoEXT-pNext-06805  
  If the `pNext` chain includes a
  [VkExportMetalSharedEventInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalSharedEventInfoEXT.html)
  structure, and its `semaphore` member is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the
  [VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html) in its `semaphore` member **must**
  have been created with
  `VK_EXPORT_METAL_OBJECT_TYPE_METAL_SHARED_EVENT_BIT_EXT` in the
  `exportObjectType` member of a
  [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalObjectCreateInfoEXT.html)
  structure in the `pNext` chain of the
  [VkSemaphoreCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreCreateInfo.html) structure in the
  [vkCreateSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateSemaphore.html) command

- <a href="#VUID-VkExportMetalObjectsInfoEXT-pNext-06806"
  id="VUID-VkExportMetalObjectsInfoEXT-pNext-06806"></a>
  VUID-VkExportMetalObjectsInfoEXT-pNext-06806  
  If the `pNext` chain includes a
  [VkExportMetalSharedEventInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalSharedEventInfoEXT.html)
  structure, and its `event` member is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the [VkEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkEvent.html) in
  its `event` member **must** have been created with
  `VK_EXPORT_METAL_OBJECT_TYPE_METAL_SHARED_EVENT_BIT_EXT` in the
  `exportObjectType` member of a
  [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalObjectCreateInfoEXT.html)
  structure in the `pNext` chain of the
  [VkEventCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkEventCreateInfo.html) structure in the
  [vkCreateEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateEvent.html) command

Valid Usage (Implicit)

- <a href="#VUID-VkExportMetalObjectsInfoEXT-sType-sType"
  id="VUID-VkExportMetalObjectsInfoEXT-sType-sType"></a>
  VUID-VkExportMetalObjectsInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_EXPORT_METAL_OBJECTS_INFO_EXT`

- <a href="#VUID-VkExportMetalObjectsInfoEXT-pNext-pNext"
  id="VUID-VkExportMetalObjectsInfoEXT-pNext-pNext"></a>
  VUID-VkExportMetalObjectsInfoEXT-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkExportMetalBufferInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalBufferInfoEXT.html),
  [VkExportMetalCommandQueueInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalCommandQueueInfoEXT.html),
  [VkExportMetalDeviceInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalDeviceInfoEXT.html),
  [VkExportMetalIOSurfaceInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalIOSurfaceInfoEXT.html),
  [VkExportMetalSharedEventInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalSharedEventInfoEXT.html),
  or [VkExportMetalTextureInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalTextureInfoEXT.html)

- <a href="#VUID-VkExportMetalObjectsInfoEXT-sType-unique"
  id="VUID-VkExportMetalObjectsInfoEXT-sType-unique"></a>
  VUID-VkExportMetalObjectsInfoEXT-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique, with the exception of structures of type
  [VkExportMetalBufferInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalBufferInfoEXT.html),
  [VkExportMetalCommandQueueInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalCommandQueueInfoEXT.html),
  [VkExportMetalIOSurfaceInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalIOSurfaceInfoEXT.html),
  [VkExportMetalSharedEventInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalSharedEventInfoEXT.html),
  or [VkExportMetalTextureInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalTextureInfoEXT.html)

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_metal_objects](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_metal_objects.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkExportMetalObjectsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkExportMetalObjectsEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkExportMetalObjectsInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
