# VkImageStencilUsageCreateInfo(3) Manual Page

## Name

VkImageStencilUsageCreateInfo - Specify separate usage flags for the
stencil aspect of a depth-stencil image



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkImageStencilUsageCreateInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_2
typedef struct VkImageStencilUsageCreateInfo {
    VkStructureType      sType;
    const void*          pNext;
    VkImageUsageFlags    stencilUsage;
} VkImageStencilUsageCreateInfo;
```

or the equivalent

``` c
// Provided by VK_EXT_separate_stencil_usage
typedef VkImageStencilUsageCreateInfo VkImageStencilUsageCreateInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `stencilUsage` is a bitmask of
  [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlagBits.html) describing the
  intended usage of the stencil aspect of the image.

## <a href="#_description" class="anchor"></a>Description

If the `pNext` chain of [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)
includes a `VkImageStencilUsageCreateInfo` structure, then that
structure includes the usage flags specific to the stencil aspect of the
image for an image with a depth-stencil format.

This structure specifies image usages which only apply to the stencil
aspect of a depth/stencil format image. When this structure is included
in the `pNext` chain of [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html), the
stencil aspect of the image **must** only be used as specified by
`stencilUsage`. When this structure is not included in the `pNext` chain
of [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html), the stencil aspect of an
image **must** only be used as specified by
[VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`usage`. Use of other
aspects of an image are unaffected by this structure.

This structure **can** also be included in the `pNext` chain of
[VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageFormatInfo2.html)
to query additional capabilities specific to image creation parameter
combinations including a separate set of usage flags for the stencil
aspect of the image using
[vkGetPhysicalDeviceImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties2.html).
When this structure is not included in the `pNext` chain of
`VkPhysicalDeviceImageFormatInfo2` then the implicit value of
`stencilUsage` matches that of
`VkPhysicalDeviceImageFormatInfo2`::`usage`.

Valid Usage

- <a href="#VUID-VkImageStencilUsageCreateInfo-stencilUsage-02539"
  id="VUID-VkImageStencilUsageCreateInfo-stencilUsage-02539"></a>
  VUID-VkImageStencilUsageCreateInfo-stencilUsage-02539  
  If `stencilUsage` includes `VK_IMAGE_USAGE_TRANSIENT_ATTACHMENT_BIT`,
  it **must** not include bits other than
  `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT` or
  `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT`

Valid Usage (Implicit)

- <a href="#VUID-VkImageStencilUsageCreateInfo-sType-sType"
  id="VUID-VkImageStencilUsageCreateInfo-sType-sType"></a>
  VUID-VkImageStencilUsageCreateInfo-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_IMAGE_STENCIL_USAGE_CREATE_INFO`

- <a href="#VUID-VkImageStencilUsageCreateInfo-stencilUsage-parameter"
  id="VUID-VkImageStencilUsageCreateInfo-stencilUsage-parameter"></a>
  VUID-VkImageStencilUsageCreateInfo-stencilUsage-parameter  
  `stencilUsage` **must** be a valid combination of
  [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlagBits.html) values

- <a
  href="#VUID-VkImageStencilUsageCreateInfo-stencilUsage-requiredbitmask"
  id="VUID-VkImageStencilUsageCreateInfo-stencilUsage-requiredbitmask"></a>
  VUID-VkImageStencilUsageCreateInfo-stencilUsage-requiredbitmask  
  `stencilUsage` **must** not be `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_separate_stencil_usage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_separate_stencil_usage.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html),
[VkImageUsageFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImageStencilUsageCreateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
