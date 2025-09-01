# VkPhysicalDeviceShaderAtomicFloat16VectorFeaturesNV(3) Manual Page

## Name

VkPhysicalDeviceShaderAtomicFloat16VectorFeaturesNV - Structure describing features supported by VK\_NV\_shader\_atomic\_float16\_vector



## [](#_c_specification)C Specification

The [VkPhysicalDeviceShaderAtomicFloat16VectorFeaturesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceShaderAtomicFloat16VectorFeaturesNV.html) structure is defined as:

```c++
// Provided by VK_NV_shader_atomic_float16_vector
typedef struct VkPhysicalDeviceShaderAtomicFloat16VectorFeaturesNV {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           shaderFloat16VectorAtomics;
} VkPhysicalDeviceShaderAtomicFloat16VectorFeaturesNV;
```

## [](#_members)Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.

## [](#_description)Description

- []()`shaderFloat16VectorAtomics` indicates whether shaders **can** perform 16-bit floating-point, 2- and 4-component vector atomic operations.

If the `VkPhysicalDeviceShaderAtomicFloat16VectorFeaturesNV` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceShaderAtomicFloat16VectorFeaturesNV`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceShaderAtomicFloat16VectorFeaturesNV-sType-sType)VUID-VkPhysicalDeviceShaderAtomicFloat16VectorFeaturesNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_ATOMIC_FLOAT16_VECTOR_FEATURES_NV`

## [](#_see_also)See Also

[VK\_NV\_shader\_atomic\_float16\_vector](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_shader_atomic_float16_vector.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceShaderAtomicFloat16VectorFeaturesNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0