# VkPhysicalDeviceImageProcessing2FeaturesQCOM(3) Manual Page

## Name

VkPhysicalDeviceImageProcessing2FeaturesQCOM - Structure describing image processing features that can be supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceImageProcessing2FeaturesQCOM` structure is defined as:

```c++
// Provided by VK_QCOM_image_processing2
typedef struct VkPhysicalDeviceImageProcessing2FeaturesQCOM {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           textureBlockMatch2;
} VkPhysicalDeviceImageProcessing2FeaturesQCOM;
```

## [](#_members)Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`textureBlockMatch2` indicates that the implementation supports shader modules that declare the `TextureBlockMatch2QCOM` capability.

## [](#_description)Description

If the `VkPhysicalDeviceImageProcessing2FeaturesQCOM` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceImageProcessing2FeaturesQCOM`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceImageProcessing2FeaturesQCOM-sType-sType)VUID-VkPhysicalDeviceImageProcessing2FeaturesQCOM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_PROCESSING_2_FEATURES_QCOM`

## [](#_see_also)See Also

[VK\_QCOM\_image\_processing2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_QCOM_image_processing2.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceImageProcessing2FeaturesQCOM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0