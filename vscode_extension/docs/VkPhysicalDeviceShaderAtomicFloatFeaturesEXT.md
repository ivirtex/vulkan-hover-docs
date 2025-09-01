# VkPhysicalDeviceShaderAtomicFloatFeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceShaderAtomicFloatFeaturesEXT - Structure describing features supported by VK\_EXT\_shader\_atomic\_float



## [](#_c_specification)C Specification

The [VkPhysicalDeviceShaderAtomicFloatFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceShaderAtomicFloatFeaturesEXT.html) structure is defined as:

```c++
// Provided by VK_EXT_shader_atomic_float
typedef struct VkPhysicalDeviceShaderAtomicFloatFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           shaderBufferFloat32Atomics;
    VkBool32           shaderBufferFloat32AtomicAdd;
    VkBool32           shaderBufferFloat64Atomics;
    VkBool32           shaderBufferFloat64AtomicAdd;
    VkBool32           shaderSharedFloat32Atomics;
    VkBool32           shaderSharedFloat32AtomicAdd;
    VkBool32           shaderSharedFloat64Atomics;
    VkBool32           shaderSharedFloat64AtomicAdd;
    VkBool32           shaderImageFloat32Atomics;
    VkBool32           shaderImageFloat32AtomicAdd;
    VkBool32           sparseImageFloat32Atomics;
    VkBool32           sparseImageFloat32AtomicAdd;
} VkPhysicalDeviceShaderAtomicFloatFeaturesEXT;
```

## [](#_members)Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.

## [](#_description)Description

- []()`shaderBufferFloat32Atomics` indicates whether shaders **can** perform 32-bit floating-point load, store and exchange atomic operations on storage buffers.
- []()`shaderBufferFloat32AtomicAdd` indicates whether shaders **can** perform 32-bit floating-point add atomic operations on storage buffers.
- []()`shaderBufferFloat64Atomics` indicates whether shaders **can** perform 64-bit floating-point load, store and exchange atomic operations on storage buffers.
- []()`shaderBufferFloat64AtomicAdd` indicates whether shaders **can** perform 64-bit floating-point add atomic operations on storage buffers.
- []()`shaderSharedFloat32Atomics` indicates whether shaders **can** perform 32-bit floating-point load, store and exchange atomic operations on shared and payload memory.
- []()`shaderSharedFloat32AtomicAdd` indicates whether shaders **can** perform 32-bit floating-point add atomic operations on shared and payload memory.
- []()`shaderSharedFloat64Atomics` indicates whether shaders **can** perform 64-bit floating-point load, store and exchange atomic operations on shared and payload memory.
- []()`shaderSharedFloat64AtomicAdd` indicates whether shaders **can** perform 64-bit floating-point add atomic operations on shared and payload memory.
- []()`shaderImageFloat32Atomics` indicates whether shaders **can** perform 32-bit floating-point load, store and exchange atomic image operations.
- []()`shaderImageFloat32AtomicAdd` indicates whether shaders **can** perform 32-bit floating-point add atomic image operations.
- []()`sparseImageFloat32Atomics` indicates whether 32-bit floating-point load, store and exchange atomic operations **can** be used on sparse images.
- []()`sparseImageFloat32AtomicAdd` indicates whether 32-bit floating-point add atomic operations **can** be used on sparse images.

If the `VkPhysicalDeviceShaderAtomicFloatFeaturesEXT` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceShaderAtomicFloatFeaturesEXT`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceShaderAtomicFloatFeaturesEXT-sType-sType)VUID-VkPhysicalDeviceShaderAtomicFloatFeaturesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_ATOMIC_FLOAT_FEATURES_EXT`

## [](#_see_also)See Also

[VK\_EXT\_shader\_atomic\_float](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_shader_atomic_float.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceShaderAtomicFloatFeaturesEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0