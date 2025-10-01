# VkPhysicalDeviceVideoEncodeRgbConversionFeaturesVALVE(3) Manual Page

## Name

VkPhysicalDeviceVideoEncodeRgbConversionFeaturesVALVE - Structure describing the video encode RGB conversion features that can be supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceVideoEncodeRgbConversionFeaturesVALVE` structure is defined as:

```c++
// Provided by VK_VALVE_video_encode_rgb_conversion
typedef struct VkPhysicalDeviceVideoEncodeRgbConversionFeaturesVALVE {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           videoEncodeRgbConversion;
} VkPhysicalDeviceVideoEncodeRgbConversionFeaturesVALVE;
```

## [](#_members)Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`videoEncodeRgbConversion` specifies that the implementation supports [video encode RGB conversion](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-rgb-conversion).

## [](#_description)Description

If the `VkPhysicalDeviceVideoEncodeRgbConversionFeaturesVALVE` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceVideoEncodeRgbConversionFeaturesVALVE`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceVideoEncodeRgbConversionFeaturesVALVE-sType-sType)VUID-VkPhysicalDeviceVideoEncodeRgbConversionFeaturesVALVE-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VIDEO_ENCODE_RGB_CONVERSION_FEATURES_VALVE`

## [](#_see_also)See Also

[VK\_VALVE\_video\_encode\_rgb\_conversion](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VALVE_video_encode_rgb_conversion.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceVideoEncodeRgbConversionFeaturesVALVE).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0