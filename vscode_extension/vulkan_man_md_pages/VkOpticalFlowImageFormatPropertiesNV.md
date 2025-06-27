# VkOpticalFlowImageFormatPropertiesNV(3) Manual Page

## Name

VkOpticalFlowImageFormatPropertiesNV - Structure describing properties
of an optical flow image format



## <a href="#_c_specification" class="anchor"></a>C Specification

The
[VkOpticalFlowImageFormatPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowImageFormatPropertiesNV.html)
structure is defined as:

``` c
// Provided by VK_NV_optical_flow
typedef struct VkOpticalFlowImageFormatPropertiesNV {
    VkStructureType    sType;
    const void*        pNext;
    VkFormat           format;
} VkOpticalFlowImageFormatPropertiesNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="opticalflow-format"></span> `format` is a
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) that specifies the format that can be used
  with the specified optical flow image usages.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkOpticalFlowImageFormatPropertiesNV-sType-sType"
  id="VUID-VkOpticalFlowImageFormatPropertiesNV-sType-sType"></a>
  VUID-VkOpticalFlowImageFormatPropertiesNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_OPTICAL_FLOW_IMAGE_FORMAT_PROPERTIES_NV`

- <a href="#VUID-VkOpticalFlowImageFormatPropertiesNV-pNext-pNext"
  id="VUID-VkOpticalFlowImageFormatPropertiesNV-pNext-pNext"></a>
  VUID-VkOpticalFlowImageFormatPropertiesNV-pNext-pNext  
  `pNext` **must** be `NULL`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_optical_flow](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_optical_flow.html),
[VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetPhysicalDeviceOpticalFlowImageFormatsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceOpticalFlowImageFormatsNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkOpticalFlowImageFormatPropertiesNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
