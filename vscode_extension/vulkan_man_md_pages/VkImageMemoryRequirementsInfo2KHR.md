# VkImageMemoryRequirementsInfo2(3) Manual Page

## Name

VkImageMemoryRequirementsInfo2 - (None)



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkImageMemoryRequirementsInfo2` structure is defined as:

``` c
// Provided by VK_VERSION_1_1
typedef struct VkImageMemoryRequirementsInfo2 {
    VkStructureType    sType;
    const void*        pNext;
    VkImage            image;
} VkImageMemoryRequirementsInfo2;
```

or the equivalent

``` c
// Provided by VK_KHR_get_memory_requirements2
typedef VkImageMemoryRequirementsInfo2 VkImageMemoryRequirementsInfo2KHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `image` is the image to query.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkImageMemoryRequirementsInfo2-image-01589"
  id="VUID-VkImageMemoryRequirementsInfo2-image-01589"></a>
  VUID-VkImageMemoryRequirementsInfo2-image-01589  
  If `image` was created with a *multi-planar* format and the
  `VK_IMAGE_CREATE_DISJOINT_BIT` flag, there **must** be a
  [VkImagePlaneMemoryRequirementsInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImagePlaneMemoryRequirementsInfo.html)
  included in the `pNext` chain of the
  [VkImageMemoryRequirementsInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageMemoryRequirementsInfo2.html)
  structure

- <a href="#VUID-VkImageMemoryRequirementsInfo2-image-02279"
  id="VUID-VkImageMemoryRequirementsInfo2-image-02279"></a>
  VUID-VkImageMemoryRequirementsInfo2-image-02279  
  If `image` was created with `VK_IMAGE_CREATE_DISJOINT_BIT` and with
  `VK_IMAGE_TILING_DRM_FORMAT_MODIFIER_EXT`, then there **must** be a
  [VkImagePlaneMemoryRequirementsInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImagePlaneMemoryRequirementsInfo.html)
  included in the `pNext` chain of the
  [VkImageMemoryRequirementsInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageMemoryRequirementsInfo2.html)
  structure

- <a href="#VUID-VkImageMemoryRequirementsInfo2-image-01590"
  id="VUID-VkImageMemoryRequirementsInfo2-image-01590"></a>
  VUID-VkImageMemoryRequirementsInfo2-image-01590  
  If `image` was not created with the `VK_IMAGE_CREATE_DISJOINT_BIT`
  flag, there **must** not be a
  [VkImagePlaneMemoryRequirementsInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImagePlaneMemoryRequirementsInfo.html)
  included in the `pNext` chain of the
  [VkImageMemoryRequirementsInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageMemoryRequirementsInfo2.html)
  structure

- <a href="#VUID-VkImageMemoryRequirementsInfo2-image-02280"
  id="VUID-VkImageMemoryRequirementsInfo2-image-02280"></a>
  VUID-VkImageMemoryRequirementsInfo2-image-02280  
  If `image` was created with a single-plane format and with any
  `tiling` other than `VK_IMAGE_TILING_DRM_FORMAT_MODIFIER_EXT`, then
  there **must** not be a
  [VkImagePlaneMemoryRequirementsInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImagePlaneMemoryRequirementsInfo.html)
  included in the `pNext` chain of the
  [VkImageMemoryRequirementsInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageMemoryRequirementsInfo2.html)
  structure

- <a href="#VUID-VkImageMemoryRequirementsInfo2-image-01897"
  id="VUID-VkImageMemoryRequirementsInfo2-image-01897"></a>
  VUID-VkImageMemoryRequirementsInfo2-image-01897  
  If `image` was created with the
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_ANDROID_HARDWARE_BUFFER_BIT_ANDROID`
  external memory handle type, then `image` **must** be bound to memory

- <a href="#VUID-VkImageMemoryRequirementsInfo2-image-08961"
  id="VUID-VkImageMemoryRequirementsInfo2-image-08961"></a>
  VUID-VkImageMemoryRequirementsInfo2-image-08961  
  If `image` was created with the
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_SCREEN_BUFFER_BIT_QNX` external memory
  handle type, then `image` **must** be bound to memory

Valid Usage (Implicit)

- <a href="#VUID-VkImageMemoryRequirementsInfo2-sType-sType"
  id="VUID-VkImageMemoryRequirementsInfo2-sType-sType"></a>
  VUID-VkImageMemoryRequirementsInfo2-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_IMAGE_MEMORY_REQUIREMENTS_INFO_2`

- <a href="#VUID-VkImageMemoryRequirementsInfo2-pNext-pNext"
  id="VUID-VkImageMemoryRequirementsInfo2-pNext-pNext"></a>
  VUID-VkImageMemoryRequirementsInfo2-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of
  [VkImagePlaneMemoryRequirementsInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImagePlaneMemoryRequirementsInfo.html)

- <a href="#VUID-VkImageMemoryRequirementsInfo2-sType-unique"
  id="VUID-VkImageMemoryRequirementsInfo2-sType-unique"></a>
  VUID-VkImageMemoryRequirementsInfo2-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkImageMemoryRequirementsInfo2-image-parameter"
  id="VUID-VkImageMemoryRequirementsInfo2-image-parameter"></a>
  VUID-VkImageMemoryRequirementsInfo2-image-parameter  
  `image` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) handle

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html), [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetImageMemoryRequirements2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageMemoryRequirements2.html),
[vkGetImageMemoryRequirements2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageMemoryRequirements2KHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImageMemoryRequirementsInfo2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
