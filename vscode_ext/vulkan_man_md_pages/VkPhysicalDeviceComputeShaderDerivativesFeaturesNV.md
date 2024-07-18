# VkPhysicalDeviceComputeShaderDerivativesFeaturesNV(3) Manual Page

## Name

VkPhysicalDeviceComputeShaderDerivativesFeaturesNV - Structure
describing compute shader derivative features that can be supported by
an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceComputeShaderDerivativesFeaturesNV` structure is
defined as:

``` c
// Provided by VK_NV_compute_shader_derivatives
typedef struct VkPhysicalDeviceComputeShaderDerivativesFeaturesNV {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           computeDerivativeGroupQuads;
    VkBool32           computeDerivativeGroupLinear;
} VkPhysicalDeviceComputeShaderDerivativesFeaturesNV;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-computeDerivativeGroupQuads"></span>
  `computeDerivativeGroupQuads` indicates that the implementation
  supports the `ComputeDerivativeGroupQuadsNV` SPIR-V capability.

- <span id="features-computeDerivativeGroupLinear"></span>
  `computeDerivativeGroupLinear` indicates that the implementation
  supports the `ComputeDerivativeGroupLinearNV` SPIR-V capability.

## <a href="#_description" class="anchor"></a>Description

See <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-scope-quad"
target="_blank" rel="noopener">Quad shader scope</a> for more
information.

If the `VkPhysicalDeviceComputeShaderDerivativesFeaturesNVfeatures`.
structure is included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceComputeShaderDerivativesFeaturesNVfeatures`. **can**
also be used in the `pNext` chain of
[VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to selectively enable
these features.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceComputeShaderDerivativesFeaturesNV-sType-sType"
  id="VUID-VkPhysicalDeviceComputeShaderDerivativesFeaturesNV-sType-sType"></a>
  VUID-VkPhysicalDeviceComputeShaderDerivativesFeaturesNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_COMPUTE_SHADER_DERIVATIVES_FEATURES_NV`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_compute_shader_derivatives](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_compute_shader_derivatives.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceComputeShaderDerivativesFeaturesNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
