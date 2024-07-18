# VkPhysicalDeviceShaderAtomicFloat2FeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceShaderAtomicFloat2FeaturesEXT - Structure describing
features supported by VK_EXT_shader_atomic_float2



## <a href="#_c_specification" class="anchor"></a>C Specification

The
[VkPhysicalDeviceShaderAtomicFloat2FeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderAtomicFloat2FeaturesEXT.html)
structure is defined as:

``` c
// Provided by VK_EXT_shader_atomic_float2
typedef struct VkPhysicalDeviceShaderAtomicFloat2FeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           shaderBufferFloat16Atomics;
    VkBool32           shaderBufferFloat16AtomicAdd;
    VkBool32           shaderBufferFloat16AtomicMinMax;
    VkBool32           shaderBufferFloat32AtomicMinMax;
    VkBool32           shaderBufferFloat64AtomicMinMax;
    VkBool32           shaderSharedFloat16Atomics;
    VkBool32           shaderSharedFloat16AtomicAdd;
    VkBool32           shaderSharedFloat16AtomicMinMax;
    VkBool32           shaderSharedFloat32AtomicMinMax;
    VkBool32           shaderSharedFloat64AtomicMinMax;
    VkBool32           shaderImageFloat32AtomicMinMax;
    VkBool32           sparseImageFloat32AtomicMinMax;
} VkPhysicalDeviceShaderAtomicFloat2FeaturesEXT;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

## <a href="#_description" class="anchor"></a>Description

- <span id="features-shaderBufferFloat16Atomics"></span>
  `shaderBufferFloat16Atomics` indicates whether shaders **can** perform
  16-bit floating-point load, store, and exchange atomic operations on
  storage buffers.

- <span id="features-shaderBufferFloat16AtomicAdd"></span>
  `shaderBufferFloat16AtomicAdd` indicates whether shaders **can**
  perform 16-bit floating-point add atomic operations on storage
  buffers.

- <span id="features-shaderBufferFloat16AtomicMinMax"></span>
  `shaderBufferFloat16AtomicMinMax` indicates whether shaders **can**
  perform 16-bit floating-point min and max atomic operations on storage
  buffers.

- <span id="features-shaderBufferFloat32AtomicMinMax"></span>
  `shaderBufferFloat32AtomicMinMax` indicates whether shaders **can**
  perform 32-bit floating-point min and max atomic operations on storage
  buffers.

- <span id="features-shaderBufferFloat64AtomicMinMax"></span>
  `shaderBufferFloat64AtomicMinMax` indicates whether shaders **can**
  perform 64-bit floating-point min and max atomic operations on storage
  buffers.

- <span id="features-shaderSharedFloat16Atomics"></span>
  `shaderSharedFloat16Atomics` indicates whether shaders **can** perform
  16-bit floating-point load, store and exchange atomic operations on
  shared and payload memory.

- <span id="features-shaderSharedFloat16AtomicAdd"></span>
  `shaderSharedFloat16AtomicAdd` indicates whether shaders **can**
  perform 16-bit floating-point add atomic operations on shared and
  payload memory.

- <span id="features-shaderSharedFloat16AtomicMinMax"></span>
  `shaderSharedFloat16AtomicMinMax` indicates whether shaders **can**
  perform 16-bit floating-point min and max atomic operations on shared
  and payload memory.

- <span id="features-shaderSharedFloat32AtomicMinMax"></span>
  `shaderSharedFloat32AtomicMinMax` indicates whether shaders **can**
  perform 32-bit floating-point min and max atomic operations on shared
  and payload memory.

- <span id="features-shaderSharedFloat64AtomicMinMax"></span>
  `shaderSharedFloat64AtomicMinMax` indicates whether shaders **can**
  perform 64-bit floating-point min and max atomic operations on shared
  and payload memory.

- <span id="features-shaderImageFloat32AtomicMinMax"></span>
  `shaderImageFloat32AtomicMinMax` indicates whether shaders **can**
  perform 32-bit floating-point min and max atomic image operations.

- <span id="features-sparseImageFloat32AtomicMinMax"></span>
  `sparseImageFloat32AtomicMinMax` indicates whether 32-bit
  floating-point min and max atomic operations **can** be used on sparse
  images.

If the `VkPhysicalDeviceShaderAtomicFloat2FeaturesEXT` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceShaderAtomicFloat2FeaturesEXT` **can** also be used in
the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceShaderAtomicFloat2FeaturesEXT-sType-sType"
  id="VUID-VkPhysicalDeviceShaderAtomicFloat2FeaturesEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceShaderAtomicFloat2FeaturesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_ATOMIC_FLOAT_2_FEATURES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_shader_atomic_float2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_shader_atomic_float2.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceShaderAtomicFloat2FeaturesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
