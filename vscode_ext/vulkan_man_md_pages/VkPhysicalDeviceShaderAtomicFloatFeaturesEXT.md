# VkPhysicalDeviceShaderAtomicFloatFeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceShaderAtomicFloatFeaturesEXT - Structure describing
features supported by VK_EXT_shader_atomic_float



## <a href="#_c_specification" class="anchor"></a>C Specification

The
[VkPhysicalDeviceShaderAtomicFloatFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderAtomicFloatFeaturesEXT.html)
structure is defined as:

``` c
// Provided by VK_EXT_shader_atomic_float
typedef struct VkPhysicalDeviceShaderAtomicFloatFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           shaderBufferFloat32Atomics;
    VkBool32           shaderBufferFloat32AtomicAdd;
    VkBool32           shaderBufferFloat64Atomics;
    VkBool32           shaderBufferFloat64AtomicAdd;
    VkBool32           shaderSharedFloat32Atomics;
    VkBool32           shaderSharedFloat32AtomicAdd;
    VkBool32           shaderSharedFloat64Atomics;
    VkBool32           shaderSharedFloat64AtomicAdd;
    VkBool32           shaderImageFloat32Atomics;
    VkBool32           shaderImageFloat32AtomicAdd;
    VkBool32           sparseImageFloat32Atomics;
    VkBool32           sparseImageFloat32AtomicAdd;
} VkPhysicalDeviceShaderAtomicFloatFeaturesEXT;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

## <a href="#_description" class="anchor"></a>Description

- <span id="features-shaderBufferFloat32Atomics"></span>
  `shaderBufferFloat32Atomics` indicates whether shaders **can** perform
  32-bit floating-point load, store and exchange atomic operations on
  storage buffers.

- <span id="features-shaderBufferFloat32AtomicAdd"></span>
  `shaderBufferFloat32AtomicAdd` indicates whether shaders **can**
  perform 32-bit floating-point add atomic operations on storage
  buffers.

- <span id="features-shaderBufferFloat64Atomics"></span>
  `shaderBufferFloat64Atomics` indicates whether shaders **can** perform
  64-bit floating-point load, store and exchange atomic operations on
  storage buffers.

- <span id="features-shaderBufferFloat64AtomicAdd"></span>
  `shaderBufferFloat64AtomicAdd` indicates whether shaders **can**
  perform 64-bit floating-point add atomic operations on storage
  buffers.

- <span id="features-shaderSharedFloat32Atomics"></span>
  `shaderSharedFloat32Atomics` indicates whether shaders **can** perform
  32-bit floating-point load, store and exchange atomic operations on
  shared and payload memory.

- <span id="features-shaderSharedFloat32AtomicAdd"></span>
  `shaderSharedFloat32AtomicAdd` indicates whether shaders **can**
  perform 32-bit floating-point add atomic operations on shared and
  payload memory.

- <span id="features-shaderSharedFloat64Atomics"></span>
  `shaderSharedFloat64Atomics` indicates whether shaders **can** perform
  64-bit floating-point load, store and exchange atomic operations on
  shared and payload memory.

- <span id="features-shaderSharedFloat64AtomicAdd"></span>
  `shaderSharedFloat64AtomicAdd` indicates whether shaders **can**
  perform 64-bit floating-point add atomic operations on shared and
  payload memory.

- <span id="features-shaderImageFloat32Atomics"></span>
  `shaderImageFloat32Atomics` indicates whether shaders **can** perform
  32-bit floating-point load, store and exchange atomic image
  operations.

- <span id="features-shaderImageFloat32AtomicAdd"></span>
  `shaderImageFloat32AtomicAdd` indicates whether shaders **can**
  perform 32-bit floating-point add atomic image operations.

- <span id="features-sparseImageFloat32Atomics"></span>
  `sparseImageFloat32Atomics` indicates whether 32-bit floating-point
  load, store and exchange atomic operations **can** be used on sparse
  images.

- <span id="features-sparseImageFloat32AtomicAdd"></span>
  `sparseImageFloat32AtomicAdd` indicates whether 32-bit floating-point
  add atomic operations **can** be used on sparse images.

If the `VkPhysicalDeviceShaderAtomicFloatFeaturesEXT` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceShaderAtomicFloatFeaturesEXT` **can** also be used in
the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceShaderAtomicFloatFeaturesEXT-sType-sType"
  id="VUID-VkPhysicalDeviceShaderAtomicFloatFeaturesEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceShaderAtomicFloatFeaturesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_ATOMIC_FLOAT_FEATURES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_shader_atomic_float](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_shader_atomic_float.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceShaderAtomicFloatFeaturesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
