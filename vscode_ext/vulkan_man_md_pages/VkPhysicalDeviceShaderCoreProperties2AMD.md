# VkPhysicalDeviceShaderCoreProperties2AMD(3) Manual Page

## Name

VkPhysicalDeviceShaderCoreProperties2AMD - Structure describing shader
core properties that can be supported by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceShaderCoreProperties2AMD` structure is defined as:

``` c
// Provided by VK_AMD_shader_core_properties2
typedef struct VkPhysicalDeviceShaderCoreProperties2AMD {
    VkStructureType                   sType;
    void*                             pNext;
    VkShaderCorePropertiesFlagsAMD    shaderCoreFeatures;
    uint32_t                          activeComputeUnitCount;
} VkPhysicalDeviceShaderCoreProperties2AMD;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-shaderCoreFeatures"></span> `shaderCoreFeatures` is
  a bitmask of
  [VkShaderCorePropertiesFlagBitsAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderCorePropertiesFlagBitsAMD.html)
  indicating the set of features supported by the shader core.

- <span id="limits-activeComputeUnitCount"></span>
  `activeComputeUnitCount` is an unsigned integer value indicating the
  number of compute units that have been enabled.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceShaderCoreProperties2AMD` structure is included
in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceShaderCoreProperties2AMD-sType-sType"
  id="VUID-VkPhysicalDeviceShaderCoreProperties2AMD-sType-sType"></a>
  VUID-VkPhysicalDeviceShaderCoreProperties2AMD-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_CORE_PROPERTIES_2_AMD`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_AMD_shader_core_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_AMD_shader_core_properties2.html),
[VkShaderCorePropertiesFlagsAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderCorePropertiesFlagsAMD.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceShaderCoreProperties2AMD"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
