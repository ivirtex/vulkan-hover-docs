# VkPhysicalDeviceShaderAtomicFloat16VectorFeaturesNV(3) Manual Page

## Name

VkPhysicalDeviceShaderAtomicFloat16VectorFeaturesNV - Structure
describing features supported by VK_NV_shader_atomic_float16_vector



## <a href="#_c_specification" class="anchor"></a>C Specification

The
[VkPhysicalDeviceShaderAtomicFloat16VectorFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderAtomicFloat16VectorFeaturesNV.html)
structure is defined as:

``` c
// Provided by VK_NV_shader_atomic_float16_vector
typedef struct VkPhysicalDeviceShaderAtomicFloat16VectorFeaturesNV {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           shaderFloat16VectorAtomics;
} VkPhysicalDeviceShaderAtomicFloat16VectorFeaturesNV;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

## <a href="#_description" class="anchor"></a>Description

- <span id="features-shaderFloat16VectorAtomics"></span>
  `shaderFloat16VectorAtomics` indicates whether shaders **can** perform
  16-bit floating-point, 2- and 4-component vector atomic operations.

If the `VkPhysicalDeviceShaderAtomicFloat16VectorFeaturesNV` structure
is included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceShaderAtomicFloat16VectorFeaturesNV` **can** also be
used in the `pNext` chain of
[VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to selectively enable
these features.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceShaderAtomicFloat16VectorFeaturesNV-sType-sType"
  id="VUID-VkPhysicalDeviceShaderAtomicFloat16VectorFeaturesNV-sType-sType"></a>
  VUID-VkPhysicalDeviceShaderAtomicFloat16VectorFeaturesNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_ATOMIC_FLOAT16_VECTOR_FEATURES_NV`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_shader_atomic_float16_vector](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_shader_atomic_float16_vector.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceShaderAtomicFloat16VectorFeaturesNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
