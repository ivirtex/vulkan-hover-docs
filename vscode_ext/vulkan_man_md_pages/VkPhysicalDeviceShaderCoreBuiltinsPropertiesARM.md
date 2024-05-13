# VkPhysicalDeviceShaderCoreBuiltinsPropertiesARM(3) Manual Page

## Name

VkPhysicalDeviceShaderCoreBuiltinsPropertiesARM - Structure describing
shader core builtins properties supported by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceShaderCoreBuiltinsPropertiesARM` structure is
defined as:

``` c
// Provided by VK_ARM_shader_core_builtins
typedef struct VkPhysicalDeviceShaderCoreBuiltinsPropertiesARM {
    VkStructureType    sType;
    void*              pNext;
    uint64_t           shaderCoreMask;
    uint32_t           shaderCoreCount;
    uint32_t           shaderWarpsPerCore;
} VkPhysicalDeviceShaderCoreBuiltinsPropertiesARM;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="limits-shaderCoreMask"></span> `shaderCoreMask` is a
  bitfield where each bit set represents the presence of a shader core
  whose ID is the bit position. The highest ID for any shader core on
  the device is the position of the most significant bit set.

- <span id="limits-shaderCoreCount"></span> `shaderCoreCount` is the
  number of shader cores on the device.

- <span id="limits-shaderWarpsPerCore"></span> `shaderWarpsPerCore` is
  the maximum number of simultaneously executing warps on a shader core.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceShaderCoreBuiltinsPropertiesARM` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceShaderCoreBuiltinsPropertiesARM-sType-sType"
  id="VUID-VkPhysicalDeviceShaderCoreBuiltinsPropertiesARM-sType-sType"></a>
  VUID-VkPhysicalDeviceShaderCoreBuiltinsPropertiesARM-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_CORE_BUILTINS_PROPERTIES_ARM`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_ARM_shader_core_builtins](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_ARM_shader_core_builtins.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceShaderCoreBuiltinsPropertiesARM"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
