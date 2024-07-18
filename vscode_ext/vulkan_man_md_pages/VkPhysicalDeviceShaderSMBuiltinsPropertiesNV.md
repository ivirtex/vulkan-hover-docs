# VkPhysicalDeviceShaderSMBuiltinsPropertiesNV(3) Manual Page

## Name

VkPhysicalDeviceShaderSMBuiltinsPropertiesNV - Structure describing
shader SM Builtins properties supported by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceShaderSMBuiltinsPropertiesNV` structure is defined
as:

``` c
// Provided by VK_NV_shader_sm_builtins
typedef struct VkPhysicalDeviceShaderSMBuiltinsPropertiesNV {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           shaderSMCount;
    uint32_t           shaderWarpsPerSM;
} VkPhysicalDeviceShaderSMBuiltinsPropertiesNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="limits-shaderSMCount"></span> `shaderSMCount` is the number
  of SMs on the device.

- <span id="limits-shaderWarpsPerSM"></span> `shaderWarpsPerSM` is the
  maximum number of simultaneously executing warps on an SM.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceShaderSMBuiltinsPropertiesNV` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceShaderSMBuiltinsPropertiesNV-sType-sType"
  id="VUID-VkPhysicalDeviceShaderSMBuiltinsPropertiesNV-sType-sType"></a>
  VUID-VkPhysicalDeviceShaderSMBuiltinsPropertiesNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_SM_BUILTINS_PROPERTIES_NV`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_shader_sm_builtins](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_shader_sm_builtins.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceShaderSMBuiltinsPropertiesNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
