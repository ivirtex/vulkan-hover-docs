# VkPhysicalDeviceShaderQuadControlFeaturesKHR(3) Manual Page

## Name

VkPhysicalDeviceShaderQuadControlFeaturesKHR - Structure describing whether quad scopes are supported by the implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceShaderQuadControlFeaturesKHR` structure is defined as:

```c++
// Provided by VK_KHR_shader_quad_control
typedef struct VkPhysicalDeviceShaderQuadControlFeaturesKHR {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           shaderQuadControl;
} VkPhysicalDeviceShaderQuadControlFeaturesKHR;
```

## [](#_members)Members

This structure describes the following features:

- []()`shaderQuadControl` indicates whether the implementation supports shaders with the `QuadControlKHR` capability.

## [](#_description)Description

If the `VkPhysicalDeviceShaderQuadControlFeaturesKHR` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceShaderQuadControlFeaturesKHR`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceShaderQuadControlFeaturesKHR-sType-sType)VUID-VkPhysicalDeviceShaderQuadControlFeaturesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_QUAD_CONTROL_FEATURES_KHR`

## [](#_see_also)See Also

[VK\_KHR\_shader\_quad\_control](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_shader_quad_control.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceShaderQuadControlFeaturesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0