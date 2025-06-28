# VkPhysicalDeviceShaderBfloat16FeaturesKHR(3) Manual Page

## Name

VkPhysicalDeviceShaderBfloat16FeaturesKHR - Structure describing bfloat16 features that can be supported by the implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceShaderBfloat16FeaturesKHR` structure is defined as:

```c++
// Provided by VK_KHR_shader_bfloat16
typedef struct VkPhysicalDeviceShaderBfloat16FeaturesKHR {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           shaderBFloat16Type;
    VkBool32           shaderBFloat16DotProduct;
    VkBool32           shaderBFloat16CooperativeMatrix;
} VkPhysicalDeviceShaderBfloat16FeaturesKHR;
```

## [](#_members)Members

This structure describes the following features:

- []()`shaderBFloat16Type` indicates whether the implementation supports shaders with the `BFloat16TypeKHR` capability.
- []()`shaderBFloat16DotProduct` indicates whether the implementation supports shaders with the `BFloat16DotProductKHR` capability.
- []()`shaderBFloat16CooperativeMatrix` indicates whether the implementation supports shaders with the `BFloat16CooperativeMatrixKHR` capability.

## [](#_description)Description

If the `VkPhysicalDeviceShaderBfloat16FeaturesKHR` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceShaderBfloat16FeaturesKHR`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceShaderBfloat16FeaturesKHR-sType-sType)VUID-VkPhysicalDeviceShaderBfloat16FeaturesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_BFLOAT16_FEATURES_KHR`

## [](#_see_also)See Also

[VK\_KHR\_shader\_bfloat16](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_shader_bfloat16.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceShaderBfloat16FeaturesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0