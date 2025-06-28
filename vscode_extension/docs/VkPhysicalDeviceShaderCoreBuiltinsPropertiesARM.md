# VkPhysicalDeviceShaderCoreBuiltinsPropertiesARM(3) Manual Page

## Name

VkPhysicalDeviceShaderCoreBuiltinsPropertiesARM - Structure describing shader core builtins properties supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceShaderCoreBuiltinsPropertiesARM` structure is defined as:

```c++
// Provided by VK_ARM_shader_core_builtins
typedef struct VkPhysicalDeviceShaderCoreBuiltinsPropertiesARM {
    VkStructureType    sType;
    void*              pNext;
    uint64_t           shaderCoreMask;
    uint32_t           shaderCoreCount;
    uint32_t           shaderWarpsPerCore;
} VkPhysicalDeviceShaderCoreBuiltinsPropertiesARM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`shaderCoreMask` is a bitfield where each bit set represents the presence of a shader core whose ID is the bit position. The highest ID for any shader core on the device is the position of the most significant bit set.
- []()`shaderCoreCount` is the number of shader cores on the device.
- []()`shaderWarpsPerCore` is the maximum number of simultaneously executing warps on a shader core.

## [](#_description)Description

If the `VkPhysicalDeviceShaderCoreBuiltinsPropertiesARM` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceShaderCoreBuiltinsPropertiesARM-sType-sType)VUID-VkPhysicalDeviceShaderCoreBuiltinsPropertiesARM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_CORE_BUILTINS_PROPERTIES_ARM`

## [](#_see_also)See Also

[VK\_ARM\_shader\_core\_builtins](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_shader_core_builtins.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceShaderCoreBuiltinsPropertiesARM)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0