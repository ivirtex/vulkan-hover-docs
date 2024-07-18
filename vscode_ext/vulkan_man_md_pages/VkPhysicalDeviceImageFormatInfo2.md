# VkPhysicalDeviceImageFormatInfo2(3) Manual Page

## Name

VkPhysicalDeviceImageFormatInfo2 - Structure specifying image creation
parameters



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceImageFormatInfo2` structure is defined as:

``` c
// Provided by VK_VERSION_1_1
typedef struct VkPhysicalDeviceImageFormatInfo2 {
    VkStructureType       sType;
    const void*           pNext;
    VkFormat              format;
    VkImageType           type;
    VkImageTiling         tiling;
    VkImageUsageFlags     usage;
    VkImageCreateFlags    flags;
} VkPhysicalDeviceImageFormatInfo2;
```

or the equivalent

``` c
// Provided by VK_KHR_get_physical_device_properties2
typedef VkPhysicalDeviceImageFormatInfo2 VkPhysicalDeviceImageFormatInfo2KHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure. The `pNext` chain of `VkPhysicalDeviceImageFormatInfo2` is
  used to provide additional image parameters to
  `vkGetPhysicalDeviceImageFormatProperties2`.

- `format` is a [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) value indicating the image
  format, corresponding to
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`format`.

- `type` is a [VkImageType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageType.html) value indicating the image
  type, corresponding to
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`imageType`.

- `tiling` is a [VkImageTiling](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageTiling.html) value indicating the
  image tiling, corresponding to
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`tiling`.

- `usage` is a bitmask of
  [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlagBits.html) indicating the
  intended usage of the image, corresponding to
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`usage`.

- `flags` is a bitmask of
  [VkImageCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateFlagBits.html) indicating
  additional parameters of the image, corresponding to
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`flags`.

## <a href="#_description" class="anchor"></a>Description

The members of `VkPhysicalDeviceImageFormatInfo2` correspond to the
arguments to
[vkGetPhysicalDeviceImageFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties.html),
with `sType` and `pNext` added for extensibility.

Valid Usage

- <a href="#VUID-VkPhysicalDeviceImageFormatInfo2-tiling-02249"
  id="VUID-VkPhysicalDeviceImageFormatInfo2-tiling-02249"></a>
  VUID-VkPhysicalDeviceImageFormatInfo2-tiling-02249  
  `tiling` **must** be `VK_IMAGE_TILING_DRM_FORMAT_MODIFIER_EXT` if and
  only if the `pNext` chain includes
  [VkPhysicalDeviceImageDrmFormatModifierInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageDrmFormatModifierInfoEXT.html)

- <a href="#VUID-VkPhysicalDeviceImageFormatInfo2-tiling-02313"
  id="VUID-VkPhysicalDeviceImageFormatInfo2-tiling-02313"></a>
  VUID-VkPhysicalDeviceImageFormatInfo2-tiling-02313  
  If `tiling` is `VK_IMAGE_TILING_DRM_FORMAT_MODIFIER_EXT` and `flags`
  contains `VK_IMAGE_CREATE_MUTABLE_FORMAT_BIT`, then the `pNext` chain
  **must** include a
  [VkImageFormatListCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatListCreateInfo.html)
  structure with non-zero `viewFormatCount`

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceImageFormatInfo2-sType-sType"
  id="VUID-VkPhysicalDeviceImageFormatInfo2-sType-sType"></a>
  VUID-VkPhysicalDeviceImageFormatInfo2-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_FORMAT_INFO_2`

- <a href="#VUID-VkPhysicalDeviceImageFormatInfo2-pNext-pNext"
  id="VUID-VkPhysicalDeviceImageFormatInfo2-pNext-pNext"></a>
  VUID-VkPhysicalDeviceImageFormatInfo2-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkImageCompressionControlEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCompressionControlEXT.html),
  [VkImageFormatListCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatListCreateInfo.html),
  [VkImageStencilUsageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageStencilUsageCreateInfo.html),
  [VkOpticalFlowImageFormatInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowImageFormatInfoNV.html),
  [VkPhysicalDeviceExternalImageFormatInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExternalImageFormatInfo.html),
  [VkPhysicalDeviceImageDrmFormatModifierInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageDrmFormatModifierInfoEXT.html),
  [VkPhysicalDeviceImageViewImageFormatInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageViewImageFormatInfoEXT.html),
  or [VkVideoProfileListInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileListInfoKHR.html)

- <a href="#VUID-VkPhysicalDeviceImageFormatInfo2-sType-unique"
  id="VUID-VkPhysicalDeviceImageFormatInfo2-sType-unique"></a>
  VUID-VkPhysicalDeviceImageFormatInfo2-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkPhysicalDeviceImageFormatInfo2-format-parameter"
  id="VUID-VkPhysicalDeviceImageFormatInfo2-format-parameter"></a>
  VUID-VkPhysicalDeviceImageFormatInfo2-format-parameter  
  `format` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) value

- <a href="#VUID-VkPhysicalDeviceImageFormatInfo2-type-parameter"
  id="VUID-VkPhysicalDeviceImageFormatInfo2-type-parameter"></a>
  VUID-VkPhysicalDeviceImageFormatInfo2-type-parameter  
  `type` **must** be a valid [VkImageType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageType.html) value

- <a href="#VUID-VkPhysicalDeviceImageFormatInfo2-tiling-parameter"
  id="VUID-VkPhysicalDeviceImageFormatInfo2-tiling-parameter"></a>
  VUID-VkPhysicalDeviceImageFormatInfo2-tiling-parameter  
  `tiling` **must** be a valid [VkImageTiling](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageTiling.html) value

- <a href="#VUID-VkPhysicalDeviceImageFormatInfo2-usage-parameter"
  id="VUID-VkPhysicalDeviceImageFormatInfo2-usage-parameter"></a>
  VUID-VkPhysicalDeviceImageFormatInfo2-usage-parameter  
  `usage` **must** be a valid combination of
  [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlagBits.html) values

- <a href="#VUID-VkPhysicalDeviceImageFormatInfo2-usage-requiredbitmask"
  id="VUID-VkPhysicalDeviceImageFormatInfo2-usage-requiredbitmask"></a>
  VUID-VkPhysicalDeviceImageFormatInfo2-usage-requiredbitmask  
  `usage` **must** not be `0`

- <a href="#VUID-VkPhysicalDeviceImageFormatInfo2-flags-parameter"
  id="VUID-VkPhysicalDeviceImageFormatInfo2-flags-parameter"></a>
  VUID-VkPhysicalDeviceImageFormatInfo2-flags-parameter  
  `flags` **must** be a valid combination of
  [VkImageCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateFlagBits.html) values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html), [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html),
[VkImageCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateFlags.html),
[VkImageTiling](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageTiling.html), [VkImageType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageType.html),
[VkImageUsageFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetPhysicalDeviceImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties2.html),
[vkGetPhysicalDeviceImageFormatProperties2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties2KHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceImageFormatInfo2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
