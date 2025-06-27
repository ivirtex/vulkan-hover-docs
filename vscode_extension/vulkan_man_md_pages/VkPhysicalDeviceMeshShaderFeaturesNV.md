# VkPhysicalDeviceMeshShaderFeaturesNV(3) Manual Page

## Name

VkPhysicalDeviceMeshShaderFeaturesNV - Structure describing mesh shading
features that can be supported by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceMeshShaderFeaturesNV` structure is defined as:

``` c
// Provided by VK_NV_mesh_shader
typedef struct VkPhysicalDeviceMeshShaderFeaturesNV {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           taskShader;
    VkBool32           meshShader;
} VkPhysicalDeviceMeshShaderFeaturesNV;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `taskShader` specifies whether task shaders are supported. If this
  feature is not enabled, the `VK_SHADER_STAGE_TASK_BIT_NV` and
  `VK_PIPELINE_STAGE_TASK_SHADER_BIT_NV` enum values **must** not be
  used.

- `meshShader` specifies whether mesh shaders are supported. If this
  feature is not enabled, the `VK_SHADER_STAGE_MESH_BIT_NV` and
  `VK_PIPELINE_STAGE_MESH_SHADER_BIT_NV` enum values **must** not be
  used.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceMeshShaderFeaturesNV` structure is included in
the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceMeshShaderFeaturesNV` **can** also be used in the
`pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceMeshShaderFeaturesNV-sType-sType"
  id="VUID-VkPhysicalDeviceMeshShaderFeaturesNV-sType-sType"></a>
  VUID-VkPhysicalDeviceMeshShaderFeaturesNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MESH_SHADER_FEATURES_NV`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_mesh_shader](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_mesh_shader.html), [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceMeshShaderFeaturesNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
