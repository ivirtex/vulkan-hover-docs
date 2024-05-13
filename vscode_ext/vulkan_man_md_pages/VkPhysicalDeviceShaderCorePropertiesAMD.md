# VkPhysicalDeviceShaderCorePropertiesAMD(3) Manual Page

## Name

VkPhysicalDeviceShaderCorePropertiesAMD - Structure describing shader
core properties that can be supported by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceShaderCorePropertiesAMD` structure is defined as:

``` c
// Provided by VK_AMD_shader_core_properties
typedef struct VkPhysicalDeviceShaderCorePropertiesAMD {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           shaderEngineCount;
    uint32_t           shaderArraysPerEngineCount;
    uint32_t           computeUnitsPerShaderArray;
    uint32_t           simdPerComputeUnit;
    uint32_t           wavefrontsPerSimd;
    uint32_t           wavefrontSize;
    uint32_t           sgprsPerSimd;
    uint32_t           minSgprAllocation;
    uint32_t           maxSgprAllocation;
    uint32_t           sgprAllocationGranularity;
    uint32_t           vgprsPerSimd;
    uint32_t           minVgprAllocation;
    uint32_t           maxVgprAllocation;
    uint32_t           vgprAllocationGranularity;
} VkPhysicalDeviceShaderCorePropertiesAMD;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="limits-shaderEngineCount"></span> `shaderEngineCount` is an
  unsigned integer value indicating the number of shader engines found
  inside the shader core of the physical device.

- <span id="limits-shaderArraysPerEngineCount"></span>
  `shaderArraysPerEngineCount` is an unsigned integer value indicating
  the number of shader arrays inside a shader engine. Each shader array
  has its own scan converter, set of compute units, and a render back
  end (color and depth attachments). Shader arrays within a shader
  engine share shader processor input (wave launcher) and shader export
  (export buffer) units. Currently, a shader engine can have one or two
  shader arrays.

- <span id="limits-computeUnitsPerShaderArray"></span>
  `computeUnitsPerShaderArray` is an unsigned integer value indicating
  the physical number of compute units within a shader array. The active
  number of compute units in a shader array **may** be lower. A compute
  unit houses a set of SIMDs along with a sequencer module and a local
  data store.

- <span id="limits-simdPerComputeUnit"></span> `simdPerComputeUnit` is
  an unsigned integer value indicating the number of SIMDs inside a
  compute unit. Each SIMD processes a single instruction at a time.

- <span id="limits-wavefrontSize"></span> `wavefrontSize` is an unsigned
  integer value indicating the maximum size of a subgroup.

- <span id="limits-sgprsPerSimd"></span> `sgprsPerSimd` is an unsigned
  integer value indicating the number of physical Scalar General-Purpose
  Registers (SGPRs) per SIMD.

- <span id="limits-minSgprAllocation"></span> `minSgprAllocation` is an
  unsigned integer value indicating the minimum number of SGPRs
  allocated for a wave.

- <span id="limits-maxSgprAllocation"></span> `maxSgprAllocation` is an
  unsigned integer value indicating the maximum number of SGPRs
  allocated for a wave.

- <span id="limits-sgprAllocationGranularity"></span>
  `sgprAllocationGranularity` is an unsigned integer value indicating
  the granularity of SGPR allocation for a wave.

- <span id="limits-vgprsPerSimd"></span> `vgprsPerSimd` is an unsigned
  integer value indicating the number of physical Vector General-Purpose
  Registers (VGPRs) per SIMD.

- <span id="limits-minVgprAllocation"></span> `minVgprAllocation` is an
  unsigned integer value indicating the minimum number of VGPRs
  allocated for a wave.

- <span id="limits-maxVgprAllocation"></span> `maxVgprAllocation` is an
  unsigned integer value indicating the maximum number of VGPRs
  allocated for a wave.

- <span id="limits-vgprAllocationGranularity"></span>
  `vgprAllocationGranularity` is an unsigned integer value indicating
  the granularity of VGPR allocation for a wave.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceShaderCorePropertiesAMD` structure is included
in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceShaderCorePropertiesAMD-sType-sType"
  id="VUID-VkPhysicalDeviceShaderCorePropertiesAMD-sType-sType"></a>
  VUID-VkPhysicalDeviceShaderCorePropertiesAMD-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_CORE_PROPERTIES_AMD`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_AMD_shader_core_properties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_AMD_shader_core_properties.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceShaderCorePropertiesAMD"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
