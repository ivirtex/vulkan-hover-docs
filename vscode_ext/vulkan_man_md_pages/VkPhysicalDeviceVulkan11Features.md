# VkPhysicalDeviceVulkan11Features(3) Manual Page

## Name

VkPhysicalDeviceVulkan11Features - Structure describing the Vulkan 1.1
features that can be supported by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceVulkan11Features` structure is defined as:

``` c
// Provided by VK_VERSION_1_2
typedef struct VkPhysicalDeviceVulkan11Features {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           storageBuffer16BitAccess;
    VkBool32           uniformAndStorageBuffer16BitAccess;
    VkBool32           storagePushConstant16;
    VkBool32           storageInputOutput16;
    VkBool32           multiview;
    VkBool32           multiviewGeometryShader;
    VkBool32           multiviewTessellationShader;
    VkBool32           variablePointersStorageBuffer;
    VkBool32           variablePointers;
    VkBool32           protectedMemory;
    VkBool32           samplerYcbcrConversion;
    VkBool32           shaderDrawParameters;
} VkPhysicalDeviceVulkan11Features;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

## <a href="#_description" class="anchor"></a>Description

- <span id="features-storageBuffer16BitAccess"></span>
  `storageBuffer16BitAccess` specifies whether objects in the
  `StorageBuffer`, `ShaderRecordBufferKHR`, or `PhysicalStorageBuffer`
  storage class with the `Block` decoration **can** have 16-bit integer
  and 16-bit floating-point members. If this feature is not enabled,
  16-bit integer or 16-bit floating-point members **must** not be used
  in such objects. This also specifies whether shader modules **can**
  declare the `StorageBuffer16BitAccess` capability.

- <span id="features-uniformAndStorageBuffer16BitAccess"></span>
  `uniformAndStorageBuffer16BitAccess` specifies whether objects in the
  `Uniform` storage class with the `Block` decoration **can** have
  16-bit integer and 16-bit floating-point members. If this feature is
  not enabled, 16-bit integer or 16-bit floating-point members **must**
  not be used in such objects. This also specifies whether shader
  modules **can** declare the `UniformAndStorageBuffer16BitAccess`
  capability.

- <span id="features-storagePushConstant16"></span>
  `storagePushConstant16` specifies whether objects in the
  `PushConstant` storage class **can** have 16-bit integer and 16-bit
  floating-point members. If this feature is not enabled, 16-bit integer
  or floating-point members **must** not be used in such objects. This
  also specifies whether shader modules **can** declare the
  `StoragePushConstant16` capability.

- <span id="features-storageInputOutput16"></span>
  `storageInputOutput16` specifies whether objects in the `Input` and
  `Output` storage classes **can** have 16-bit integer and 16-bit
  floating-point members. If this feature is not enabled, 16-bit integer
  or 16-bit floating-point members **must** not be used in such objects.
  This also specifies whether shader modules **can** declare the
  `StorageInputOutput16` capability.

- <span id="features-multiview"></span> `multiview` specifies whether
  the implementation supports multiview rendering within a render pass.
  If this feature is not enabled, the view mask of each subpass **must**
  always be zero.

- <span id="features-multiview-gs"></span> `multiviewGeometryShader`
  specifies whether the implementation supports multiview rendering
  within a render pass, with [geometry shaders](#geometry). If this
  feature is not enabled, then a pipeline compiled against a subpass
  with a non-zero view mask **must** not include a geometry shader.

- <span id="features-multiview-tess"></span>
  `multiviewTessellationShader` specifies whether the implementation
  supports multiview rendering within a render pass, with [tessellation
  shaders](#tessellation). If this feature is not enabled, then a
  pipeline compiled against a subpass with a non-zero view mask **must**
  not include any tessellation shaders.

- <span id="features-variablePointersStorageBuffer"></span>
  `variablePointersStorageBuffer` specifies whether the implementation
  supports the SPIR-V `VariablePointersStorageBuffer` capability. When
  this feature is not enabled, shader modules **must** not declare the
  `SPV_KHR_variable_pointers` extension or the
  `VariablePointersStorageBuffer` capability.

- <span id="features-variablePointers"></span> `variablePointers`
  specifies whether the implementation supports the SPIR-V
  `VariablePointers` capability. When this feature is not enabled,
  shader modules **must** not declare the `VariablePointers` capability.

- <span id="features-protectedMemory"></span> `protectedMemory`
  specifies whether [protected memory](#memory-protected-memory) is
  supported.

- <span id="features-samplerYcbcrConversion"></span>
  `samplerYcbcrConversion` specifies whether the implementation supports
  [sampler Y′C<sub>B</sub>C<sub>R</sub>
  conversion](#samplers-YCbCr-conversion). If `samplerYcbcrConversion`
  is `VK_FALSE`, sampler Y′C<sub>B</sub>C<sub>R</sub> conversion is not
  supported, and samplers using sampler Y′C<sub>B</sub>C<sub>R</sub>
  conversion **must** not be used.

- <span id="features-shaderDrawParameters"></span>
  `shaderDrawParameters` specifies whether the implementation supports
  the SPIR-V `DrawParameters` capability. When this feature is not
  enabled, shader modules **must** not declare the
  `SPV_KHR_shader_draw_parameters` extension or the `DrawParameters`
  capability.

If the `VkPhysicalDeviceVulkan11Features` structure is included in the
`pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceVulkan11Features` **can** also be used in the `pNext`
chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to selectively
enable these features.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceVulkan11Features-sType-sType"
  id="VUID-VkPhysicalDeviceVulkan11Features-sType-sType"></a>
  VUID-VkPhysicalDeviceVulkan11Features-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_1_1_FEATURES`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html), [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceVulkan11Features"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
