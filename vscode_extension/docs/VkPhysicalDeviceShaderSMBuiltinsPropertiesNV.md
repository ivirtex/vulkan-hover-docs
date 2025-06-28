# VkPhysicalDeviceShaderSMBuiltinsPropertiesNV(3) Manual Page

## Name

VkPhysicalDeviceShaderSMBuiltinsPropertiesNV - Structure describing shader SM Builtins properties supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceShaderSMBuiltinsPropertiesNV` structure is defined as:

```c++
// Provided by VK_NV_shader_sm_builtins
typedef struct VkPhysicalDeviceShaderSMBuiltinsPropertiesNV {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           shaderSMCount;
    uint32_t           shaderWarpsPerSM;
} VkPhysicalDeviceShaderSMBuiltinsPropertiesNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`shaderSMCount` is the number of SMs on the device.
- []()`shaderWarpsPerSM` is the maximum number of simultaneously executing warps on an SM.

## [](#_description)Description

If the `VkPhysicalDeviceShaderSMBuiltinsPropertiesNV` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceShaderSMBuiltinsPropertiesNV-sType-sType)VUID-VkPhysicalDeviceShaderSMBuiltinsPropertiesNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_SM_BUILTINS_PROPERTIES_NV`

## [](#_see_also)See Also

[VK\_NV\_shader\_sm\_builtins](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_shader_sm_builtins.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceShaderSMBuiltinsPropertiesNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0