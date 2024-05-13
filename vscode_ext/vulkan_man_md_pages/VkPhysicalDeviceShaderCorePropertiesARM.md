# VkPhysicalDeviceShaderCorePropertiesARM(3) Manual Page

## Name

VkPhysicalDeviceShaderCorePropertiesARM - Structure describing shader
core properties that can be supported by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceShaderCorePropertiesARM` structure is defined as:

``` c
// Provided by VK_ARM_shader_core_properties
typedef struct VkPhysicalDeviceShaderCorePropertiesARM {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           pixelRate;
    uint32_t           texelRate;
    uint32_t           fmaRate;
} VkPhysicalDeviceShaderCorePropertiesARM;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `pixelRate` is an unsigned integer value indicating the maximum number
  of pixels output per clock per shader core.

- `texelRate` is an unsigned integer value indicating the maximum number
  of texels per clock per shader core.

- `fmaRate` is an unsigned integer value indicating the maximum number
  of single-precision fused multiply-add operations per clock per shader
  core.

## <a href="#_description" class="anchor"></a>Description

If a throughput rate cannot be determined on the physical device, the
value `0` will be returned for that rate.

If the `VkPhysicalDeviceShaderCorePropertiesARM` structure is included
in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceShaderCorePropertiesARM-sType-sType"
  id="VUID-VkPhysicalDeviceShaderCorePropertiesARM-sType-sType"></a>
  VUID-VkPhysicalDeviceShaderCorePropertiesARM-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_CORE_PROPERTIES_ARM`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_ARM_shader_core_properties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_ARM_shader_core_properties.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceShaderCorePropertiesARM"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
