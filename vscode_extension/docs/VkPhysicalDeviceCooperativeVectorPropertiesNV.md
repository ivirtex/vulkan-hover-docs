# VkPhysicalDeviceCooperativeVectorPropertiesNV(3) Manual Page

## Name

VkPhysicalDeviceCooperativeVectorPropertiesNV - Structure describing cooperative vector properties supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceCooperativeVectorPropertiesNV` structure is defined as:

```c++
// Provided by VK_NV_cooperative_vector
typedef struct VkPhysicalDeviceCooperativeVectorPropertiesNV {
    VkStructureType       sType;
    void*                 pNext;
    VkShaderStageFlags    cooperativeVectorSupportedStages;
    VkBool32              cooperativeVectorTrainingFloat16Accumulation;
    VkBool32              cooperativeVectorTrainingFloat32Accumulation;
    uint32_t              maxCooperativeVectorComponents;
} VkPhysicalDeviceCooperativeVectorPropertiesNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`cooperativeVectorSupportedStages` is a bitfield of [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlagBits.html) describing the shader stages that cooperative vector instructions are supported in. `cooperativeVectorSupportedStages` will have the `VK_SHADER_STAGE_COMPUTE_BIT` bit set if any of the physical device’s queues support `VK_QUEUE_COMPUTE_BIT`.
- []()`cooperativeVectorTrainingFloat16Accumulation` is `VK_TRUE` if the implementation supports cooperative vector training functions accumulating 16-bit floating-point results.
- []()`cooperativeVectorTrainingFloat32Accumulation` is `VK_TRUE` if the implementation supports cooperative vector training functions accumulating 32-bit floating-point results.
- []()`maxCooperativeVectorComponents` indicates the maximum number of components that **can** be in a cooperative vector.

## [](#_description)Description

If the `VkPhysicalDeviceCooperativeVectorPropertiesNV` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceCooperativeVectorPropertiesNV-sType-sType)VUID-VkPhysicalDeviceCooperativeVectorPropertiesNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_COOPERATIVE_VECTOR_PROPERTIES_NV`

## [](#_see_also)See Also

[VK\_NV\_cooperative\_vector](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_cooperative_vector.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkShaderStageFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceCooperativeVectorPropertiesNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0