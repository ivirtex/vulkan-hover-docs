# VkPhysicalDeviceComputeShaderDerivativesPropertiesKHR(3) Manual Page

## Name

VkPhysicalDeviceComputeShaderDerivativesPropertiesKHR - Structure describing compute shader derivative operations supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceComputeShaderDerivativesPropertiesKHR` structure is defined as:

```c++
// Provided by VK_KHR_compute_shader_derivatives
typedef struct VkPhysicalDeviceComputeShaderDerivativesPropertiesKHR {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           meshAndTaskShaderDerivatives;
} VkPhysicalDeviceComputeShaderDerivativesPropertiesKHR;
```

## [](#_members)Members

The members of the `VkPhysicalDeviceComputeShaderDerivativesPropertiesKHR` structure describe the following:

## [](#_description)Description

- []()`meshAndTaskShaderDerivatives` indicates whether the mesh and task shader stages support the `ComputeDerivativeGroupQuadsKHR` and `ComputeDerivativeGroupLinearKHR` SPIR-V capabilities.

If the `VkPhysicalDeviceComputeShaderDerivativesPropertiesKHR` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceComputeShaderDerivativesPropertiesKHR-sType-sType)VUID-VkPhysicalDeviceComputeShaderDerivativesPropertiesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_COMPUTE_SHADER_DERIVATIVES_PROPERTIES_KHR`

## [](#_see_also)See Also

[VK\_KHR\_compute\_shader\_derivatives](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_compute_shader_derivatives.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceComputeShaderDerivativesPropertiesKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0