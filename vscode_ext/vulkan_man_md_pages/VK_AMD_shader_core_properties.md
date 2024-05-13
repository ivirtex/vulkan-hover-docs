# VK_AMD_shader_core_properties(3) Manual Page

## Name

VK_AMD_shader_core_properties - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

186

## <a href="#_revision" class="anchor"></a>Revision

2

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_contact" class="anchor"></a>Contact

- Martin Dinkov <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_AMD_shader_core_properties%5D%20@mdinkov%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_AMD_shader_core_properties%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>mdinkov</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2019-06-25

**IP Status**  
No known IP claims.

**Contributors**  
- Martin Dinkov, AMD

- Matthaeus G. Chajdas, AMD

## <a href="#_description" class="anchor"></a>Description

This extension exposes shader core properties for a target physical
device through the
[`VK_KHR_get_physical_device_properties2`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)
extension. Please refer to the example below for proper usage.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceShaderCorePropertiesAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderCorePropertiesAMD.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_AMD_SHADER_CORE_PROPERTIES_EXTENSION_NAME`

- `VK_AMD_SHADER_CORE_PROPERTIES_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_CORE_PROPERTIES_AMD`

## <a href="#_examples" class="anchor"></a>Examples

This example retrieves the shader core properties for a physical device.

``` c
extern VkInstance       instance;

PFN_vkGetPhysicalDeviceProperties2 pfnVkGetPhysicalDeviceProperties2 =
    reinterpret_cast<PFN_vkGetPhysicalDeviceProperties2>
    (vkGetInstanceProcAddr(instance, "vkGetPhysicalDeviceProperties2") );

VkPhysicalDeviceProperties2             general_props;
VkPhysicalDeviceShaderCorePropertiesAMD shader_core_properties;

shader_core_properties.pNext = nullptr;
shader_core_properties.sType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_CORE_PROPERTIES_AMD;

general_props.pNext = &shader_core_properties;
general_props.sType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2;

// After this call, shader_core_properties has been populated
pfnVkGetPhysicalDeviceProperties2(device, &general_props);

printf("Number of shader engines: %d\n",
    m_shader_core_properties.shader_engine_count =
    shader_core_properties.shaderEngineCount;
printf("Number of shader arrays: %d\n",
    m_shader_core_properties.shader_arrays_per_engine_count =
    shader_core_properties.shaderArraysPerEngineCount;
printf("Number of CUs per shader array: %d\n",
    m_shader_core_properties.compute_units_per_shader_array =
    shader_core_properties.computeUnitsPerShaderArray;
printf("Number of SIMDs per compute unit: %d\n",
    m_shader_core_properties.simd_per_compute_unit =
    shader_core_properties.simdPerComputeUnit;
printf("Number of wavefront slots in each SIMD: %d\n",
    m_shader_core_properties.wavefronts_per_simd =
    shader_core_properties.wavefrontsPerSimd;
printf("Number of threads per wavefront: %d\n",
    m_shader_core_properties.wavefront_size =
    shader_core_properties.wavefrontSize;
printf("Number of physical SGPRs per SIMD: %d\n",
    m_shader_core_properties.sgprs_per_simd =
    shader_core_properties.sgprsPerSimd;
printf("Minimum number of SGPRs that can be allocated by a wave: %d\n",
    m_shader_core_properties.min_sgpr_allocation =
    shader_core_properties.minSgprAllocation;
printf("Number of available SGPRs: %d\n",
    m_shader_core_properties.max_sgpr_allocation =
    shader_core_properties.maxSgprAllocation;
printf("SGPRs are allocated in groups of this size: %d\n",
    m_shader_core_properties.sgpr_allocation_granularity =
    shader_core_properties.sgprAllocationGranularity;
printf("Number of physical VGPRs per SIMD: %d\n",
    m_shader_core_properties.vgprs_per_simd =
    shader_core_properties.vgprsPerSimd;
printf("Minimum number of VGPRs that can be allocated by a wave: %d\n",
    m_shader_core_properties.min_vgpr_allocation =
    shader_core_properties.minVgprAllocation;
printf("Number of available VGPRs: %d\n",
    m_shader_core_properties.max_vgpr_allocation =
    shader_core_properties.maxVgprAllocation;
printf("VGPRs are allocated in groups of this size: %d\n",
    m_shader_core_properties.vgpr_allocation_granularity =
    shader_core_properties.vgprAllocationGranularity;
```

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 2, 2019-06-25 (Matthaeus G. Chajdas)

  - Clarified the meaning of a few fields.

- Revision 1, 2018-02-15 (Martin Dinkov)

  - Initial draft.

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_AMD_shader_core_properties"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
