# VkVideoEncodeQualityLevelPropertiesKHR(3) Manual Page

## Name

VkVideoEncodeQualityLevelPropertiesKHR - Structure describing the video
encode quality level properties



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkVideoEncodeQualityLevelPropertiesKHR` structure is defined as:

``` c
// Provided by VK_KHR_video_encode_queue
typedef struct VkVideoEncodeQualityLevelPropertiesKHR {
    VkStructureType                            sType;
    void*                                      pNext;
    VkVideoEncodeRateControlModeFlagBitsKHR    preferredRateControlMode;
    uint32_t                                   preferredRateControlLayerCount;
} VkVideoEncodeQualityLevelPropertiesKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `preferredRateControlMode` is a
  [VkVideoEncodeRateControlModeFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeRateControlModeFlagBitsKHR.html)
  value indicating the preferred <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-rate-control-modes"
  target="_blank" rel="noopener">rate control mode</a> to use with the
  video encode quality level.

- `preferredRateControlLayerCount` indicates the preferred number of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-rate-control-layers"
  target="_blank" rel="noopener">rate control layers</a> to use with the
  video encode quality level.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkVideoEncodeQualityLevelPropertiesKHR-sType-sType"
  id="VUID-VkVideoEncodeQualityLevelPropertiesKHR-sType-sType"></a>
  VUID-VkVideoEncodeQualityLevelPropertiesKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_VIDEO_ENCODE_QUALITY_LEVEL_PROPERTIES_KHR`

- <a href="#VUID-VkVideoEncodeQualityLevelPropertiesKHR-pNext-pNext"
  id="VUID-VkVideoEncodeQualityLevelPropertiesKHR-pNext-pNext"></a>
  VUID-VkVideoEncodeQualityLevelPropertiesKHR-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkVideoEncodeH264QualityLevelPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264QualityLevelPropertiesKHR.html)
  or
  [VkVideoEncodeH265QualityLevelPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265QualityLevelPropertiesKHR.html)

- <a href="#VUID-VkVideoEncodeQualityLevelPropertiesKHR-sType-unique"
  id="VUID-VkVideoEncodeQualityLevelPropertiesKHR-sType-unique"></a>
  VUID-VkVideoEncodeQualityLevelPropertiesKHR-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_encode_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_encode_queue.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkVideoEncodeRateControlModeFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeRateControlModeFlagBitsKHR.html),
[vkGetPhysicalDeviceVideoEncodeQualityLevelPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoEncodeQualityLevelPropertiesKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoEncodeQualityLevelPropertiesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
