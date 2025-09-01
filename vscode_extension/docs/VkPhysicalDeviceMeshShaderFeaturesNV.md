# VkPhysicalDeviceMeshShaderFeaturesNV(3) Manual Page

## Name

VkPhysicalDeviceMeshShaderFeaturesNV - Structure describing mesh shading features that can be supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceMeshShaderFeaturesNV` structure is defined as:

```c++
// Provided by VK_NV_mesh_shader
typedef struct VkPhysicalDeviceMeshShaderFeaturesNV {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           taskShader;
    VkBool32           meshShader;
} VkPhysicalDeviceMeshShaderFeaturesNV;
```

## [](#_members)Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`taskShader` specifies whether task shaders are supported. If this feature is not enabled, the `VK_SHADER_STAGE_TASK_BIT_NV` and `VK_PIPELINE_STAGE_TASK_SHADER_BIT_NV` enum values **must** not be used.
- []()`meshShader` specifies whether mesh shaders are supported. If this feature is not enabled, the `VK_SHADER_STAGE_MESH_BIT_NV` and `VK_PIPELINE_STAGE_MESH_SHADER_BIT_NV` enum values **must** not be used.

## [](#_description)Description

If the `VkPhysicalDeviceMeshShaderFeaturesNV` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceMeshShaderFeaturesNV`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceMeshShaderFeaturesNV-sType-sType)VUID-VkPhysicalDeviceMeshShaderFeaturesNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MESH_SHADER_FEATURES_NV`

## [](#_see_also)See Also

[VK\_NV\_mesh\_shader](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_mesh_shader.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceMeshShaderFeaturesNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0