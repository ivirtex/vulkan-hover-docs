# VkPhysicalDeviceShaderCorePropertiesARM(3) Manual Page

## Name

VkPhysicalDeviceShaderCorePropertiesARM - Structure describing shader core properties that can be supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceShaderCorePropertiesARM` structure is defined as:

```c++
// Provided by VK_ARM_shader_core_properties
typedef struct VkPhysicalDeviceShaderCorePropertiesARM {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           pixelRate;
    uint32_t           texelRate;
    uint32_t           fmaRate;
} VkPhysicalDeviceShaderCorePropertiesARM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `pixelRate` is an unsigned integer value indicating the maximum number of pixels output per clock per shader core.
- `texelRate` is an unsigned integer value indicating the maximum number of texels per clock per shader core.
- `fmaRate` is an unsigned integer value indicating the maximum number of single-precision fused multiply-add operations per clock per shader core.

## [](#_description)Description

If a throughput rate cannot be determined on the physical device, the value `0` will be returned for that rate.

If the `VkPhysicalDeviceShaderCorePropertiesARM` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceShaderCorePropertiesARM-sType-sType)VUID-VkPhysicalDeviceShaderCorePropertiesARM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_CORE_PROPERTIES_ARM`

## [](#_see_also)See Also

[VK\_ARM\_shader\_core\_properties](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_shader_core_properties.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceShaderCorePropertiesARM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0