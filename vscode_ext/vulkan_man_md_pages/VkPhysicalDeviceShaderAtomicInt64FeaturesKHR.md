# VkPhysicalDeviceShaderAtomicInt64Features(3) Manual Page

## Name

VkPhysicalDeviceShaderAtomicInt64Features - Structure describing
features supported by VK_KHR_shader_atomic_int64



## <a href="#_c_specification" class="anchor"></a>C Specification

The
[VkPhysicalDeviceShaderAtomicInt64Features](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderAtomicInt64Features.html)
structure is defined as:

``` c
// Provided by VK_VERSION_1_2
typedef struct VkPhysicalDeviceShaderAtomicInt64Features {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           shaderBufferInt64Atomics;
    VkBool32           shaderSharedInt64Atomics;
} VkPhysicalDeviceShaderAtomicInt64Features;
```

or the equivalent

``` c
// Provided by VK_KHR_shader_atomic_int64
typedef VkPhysicalDeviceShaderAtomicInt64Features VkPhysicalDeviceShaderAtomicInt64FeaturesKHR;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

## <a href="#_description" class="anchor"></a>Description

- <span id="extension-features-shaderBufferInt64Atomics"></span>
  `shaderBufferInt64Atomics` indicates whether shaders **can** perform
  64-bit unsigned and signed integer atomic operations on buffers.

- <span id="extension-features-shaderSharedInt64Atomics"></span>
  `shaderSharedInt64Atomics` indicates whether shaders **can** perform
  64-bit unsigned and signed integer atomic operations on shared and
  payload memory.

If the `VkPhysicalDeviceShaderAtomicInt64Features` structure is included
in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceShaderAtomicInt64Features` **can** also be used in the
`pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceShaderAtomicInt64Features-sType-sType"
  id="VUID-VkPhysicalDeviceShaderAtomicInt64Features-sType-sType"></a>
  VUID-VkPhysicalDeviceShaderAtomicInt64Features-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_ATOMIC_INT64_FEATURES`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_shader_atomic_int64](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_shader_atomic_int64.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html), [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceShaderAtomicInt64Features"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
