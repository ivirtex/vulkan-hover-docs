# VkPhysicalDeviceImageProcessingFeaturesQCOM(3) Manual Page

## Name

VkPhysicalDeviceImageProcessingFeaturesQCOM - Structure describing image
processing features that can be supported by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceImageProcessingFeaturesQCOM` structure is defined
as:

``` c
// Provided by VK_QCOM_image_processing
typedef struct VkPhysicalDeviceImageProcessingFeaturesQCOM {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           textureSampleWeighted;
    VkBool32           textureBoxFilter;
    VkBool32           textureBlockMatch;
} VkPhysicalDeviceImageProcessingFeaturesQCOM;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-textureSampleWeighted"></span>
  `textureSampleWeighted` indicates that the implementation supports
  shader modules that declare the `TextureSampleWeightedQCOM`
  capability.

- <span id="features-textureBoxFilter"></span> `textureBoxFilter`
  indicates that the implementation supports shader modules that declare
  the `TextureBoxFilterQCOM` capability.

- <span id="features-textureBlockMatch"></span> `textureBlockMatch`
  indicates that the implementation supports shader modules that declare
  the `TextureBlockMatchQCOM` capability.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceImageProcessingFeaturesQCOM` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceImageProcessingFeaturesQCOM` **can** also be used in
the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceImageProcessingFeaturesQCOM-sType-sType"
  id="VUID-VkPhysicalDeviceImageProcessingFeaturesQCOM-sType-sType"></a>
  VUID-VkPhysicalDeviceImageProcessingFeaturesQCOM-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_PROCESSING_FEATURES_QCOM`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_QCOM_image_processing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_QCOM_image_processing.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceImageProcessingFeaturesQCOM"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
