# VkPhysicalDeviceShaderCoreProperties2AMD(3) Manual Page

## Name

VkPhysicalDeviceShaderCoreProperties2AMD - Structure describing shader core properties that can be supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceShaderCoreProperties2AMD` structure is defined as:

```c++
// Provided by VK_AMD_shader_core_properties2
typedef struct VkPhysicalDeviceShaderCoreProperties2AMD {
    VkStructureType                   sType;
    void*                             pNext;
    VkShaderCorePropertiesFlagsAMD    shaderCoreFeatures;
    uint32_t                          activeComputeUnitCount;
} VkPhysicalDeviceShaderCoreProperties2AMD;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`shaderCoreFeatures` is a bitmask of [VkShaderCorePropertiesFlagBitsAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderCorePropertiesFlagBitsAMD.html) indicating the set of features supported by the shader core.
- []()`activeComputeUnitCount` is an unsigned integer value indicating the number of compute units that have been enabled.

## [](#_description)Description

If the `VkPhysicalDeviceShaderCoreProperties2AMD` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceShaderCoreProperties2AMD-sType-sType)VUID-VkPhysicalDeviceShaderCoreProperties2AMD-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_CORE_PROPERTIES_2_AMD`

## [](#_see_also)See Also

[VK\_AMD\_shader\_core\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_AMD_shader_core_properties2.html), [VkShaderCorePropertiesFlagsAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderCorePropertiesFlagsAMD.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceShaderCoreProperties2AMD)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0