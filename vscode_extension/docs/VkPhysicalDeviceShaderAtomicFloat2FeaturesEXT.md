# VkPhysicalDeviceShaderAtomicFloat2FeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceShaderAtomicFloat2FeaturesEXT - Structure describing features supported by VK\_EXT\_shader\_atomic\_float2



## [](#_c_specification)C Specification

The [VkPhysicalDeviceShaderAtomicFloat2FeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceShaderAtomicFloat2FeaturesEXT.html) structure is defined as:

```c++
// Provided by VK_EXT_shader_atomic_float2
typedef struct VkPhysicalDeviceShaderAtomicFloat2FeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           shaderBufferFloat16Atomics;
    VkBool32           shaderBufferFloat16AtomicAdd;
    VkBool32           shaderBufferFloat16AtomicMinMax;
    VkBool32           shaderBufferFloat32AtomicMinMax;
    VkBool32           shaderBufferFloat64AtomicMinMax;
    VkBool32           shaderSharedFloat16Atomics;
    VkBool32           shaderSharedFloat16AtomicAdd;
    VkBool32           shaderSharedFloat16AtomicMinMax;
    VkBool32           shaderSharedFloat32AtomicMinMax;
    VkBool32           shaderSharedFloat64AtomicMinMax;
    VkBool32           shaderImageFloat32AtomicMinMax;
    VkBool32           sparseImageFloat32AtomicMinMax;
} VkPhysicalDeviceShaderAtomicFloat2FeaturesEXT;
```

## [](#_members)Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.

## [](#_description)Description

- []()`shaderBufferFloat16Atomics` indicates whether shaders **can** perform 16-bit floating-point load, store, and exchange atomic operations on storage buffers.
- []()`shaderBufferFloat16AtomicAdd` indicates whether shaders **can** perform 16-bit floating-point add atomic operations on storage buffers.
- []()`shaderBufferFloat16AtomicMinMax` indicates whether shaders **can** perform 16-bit floating-point min and max atomic operations on storage buffers.
- []()`shaderBufferFloat32AtomicMinMax` indicates whether shaders **can** perform 32-bit floating-point min and max atomic operations on storage buffers.
- []()`shaderBufferFloat64AtomicMinMax` indicates whether shaders **can** perform 64-bit floating-point min and max atomic operations on storage buffers.
- []()`shaderSharedFloat16Atomics` indicates whether shaders **can** perform 16-bit floating-point load, store and exchange atomic operations on shared and payload memory.
- []()`shaderSharedFloat16AtomicAdd` indicates whether shaders **can** perform 16-bit floating-point add atomic operations on shared and payload memory.
- []()`shaderSharedFloat16AtomicMinMax` indicates whether shaders **can** perform 16-bit floating-point min and max atomic operations on shared and payload memory.
- []()`shaderSharedFloat32AtomicMinMax` indicates whether shaders **can** perform 32-bit floating-point min and max atomic operations on shared and payload memory.
- []()`shaderSharedFloat64AtomicMinMax` indicates whether shaders **can** perform 64-bit floating-point min and max atomic operations on shared and payload memory.
- []()`shaderImageFloat32AtomicMinMax` indicates whether shaders **can** perform 32-bit floating-point min and max atomic image operations.
- []()`sparseImageFloat32AtomicMinMax` indicates whether 32-bit floating-point min and max atomic operations **can** be used on sparse images.

If the `VkPhysicalDeviceShaderAtomicFloat2FeaturesEXT` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceShaderAtomicFloat2FeaturesEXT`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceShaderAtomicFloat2FeaturesEXT-sType-sType)VUID-VkPhysicalDeviceShaderAtomicFloat2FeaturesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_ATOMIC_FLOAT_2_FEATURES_EXT`

## [](#_see_also)See Also

[VK\_EXT\_shader\_atomic\_float2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_shader_atomic_float2.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceShaderAtomicFloat2FeaturesEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0