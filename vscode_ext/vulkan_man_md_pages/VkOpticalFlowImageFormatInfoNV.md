# VkOpticalFlowImageFormatInfoNV(3) Manual Page

## Name

VkOpticalFlowImageFormatInfoNV - Structure describing optical flow image
format info



## <a href="#_c_specification" class="anchor"></a>C Specification

The
[VkOpticalFlowImageFormatInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowImageFormatInfoNV.html)
structure is defined as:

``` c
// Provided by VK_NV_optical_flow
typedef struct VkOpticalFlowImageFormatInfoNV {
    VkStructureType              sType;
    const void*                  pNext;
    VkOpticalFlowUsageFlagsNV    usage;
} VkOpticalFlowImageFormatInfoNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="opticalflow-usage"></span> `usage` is a bitmask of
  [VkOpticalFlowUsageFlagBitsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowUsageFlagBitsNV.html)
  describing the intended optical flow usage of the image.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkOpticalFlowImageFormatInfoNV-sType-sType"
  id="VUID-VkOpticalFlowImageFormatInfoNV-sType-sType"></a>
  VUID-VkOpticalFlowImageFormatInfoNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_OPTICAL_FLOW_IMAGE_FORMAT_INFO_NV`

- <a href="#VUID-VkOpticalFlowImageFormatInfoNV-usage-parameter"
  id="VUID-VkOpticalFlowImageFormatInfoNV-usage-parameter"></a>
  VUID-VkOpticalFlowImageFormatInfoNV-usage-parameter  
  `usage` **must** be a valid combination of
  [VkOpticalFlowUsageFlagBitsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowUsageFlagBitsNV.html)
  values

- <a href="#VUID-VkOpticalFlowImageFormatInfoNV-usage-requiredbitmask"
  id="VUID-VkOpticalFlowImageFormatInfoNV-usage-requiredbitmask"></a>
  VUID-VkOpticalFlowImageFormatInfoNV-usage-requiredbitmask  
  `usage` **must** not be `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_optical_flow](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_optical_flow.html),
[VkOpticalFlowUsageFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowUsageFlagsNV.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetPhysicalDeviceOpticalFlowImageFormatsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceOpticalFlowImageFormatsNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkOpticalFlowImageFormatInfoNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
