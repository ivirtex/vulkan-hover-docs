# VkPhysicalDeviceShaderFloat16Int8Features(3) Manual Page

## Name

VkPhysicalDeviceShaderFloat16Int8Features - Structure describing
features supported by VK_KHR_shader_float16_int8



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceShaderFloat16Int8Features` structure is defined as:

``` c
// Provided by VK_VERSION_1_2
typedef struct VkPhysicalDeviceShaderFloat16Int8Features {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           shaderFloat16;
    VkBool32           shaderInt8;
} VkPhysicalDeviceShaderFloat16Int8Features;
```

or the equivalent

``` c
// Provided by VK_KHR_shader_float16_int8
typedef VkPhysicalDeviceShaderFloat16Int8Features VkPhysicalDeviceShaderFloat16Int8FeaturesKHR;
```

``` c
// Provided by VK_KHR_shader_float16_int8
typedef VkPhysicalDeviceShaderFloat16Int8Features VkPhysicalDeviceFloat16Int8FeaturesKHR;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

## <a href="#_description" class="anchor"></a>Description

- <span id="extension-features-shaderFloat16"></span> `shaderFloat16`
  indicates whether 16-bit floats (halfs) are supported in shader code.
  This also indicates whether shader modules **can** declare the
  `Float16` capability. However, this only enables a subset of the
  storage classes that SPIR-V allows for the `Float16` SPIR-V
  capability: Declaring and using 16-bit floats in the `Private`,
  `Workgroup` (for non-Block variables), and `Function` storage classes
  is enabled, while declaring them in the interface storage classes
  (e.g., `UniformConstant`, `Uniform`, `StorageBuffer`, `Input`,
  `Output`, and `PushConstant`) is not enabled.

- <span id="extension-features-shaderInt8"></span> `shaderInt8`
  indicates whether 8-bit integers (signed and unsigned) are supported
  in shader code. This also indicates whether shader modules **can**
  declare the `Int8` capability. However, this only enables a subset of
  the storage classes that SPIR-V allows for the `Int8` SPIR-V
  capability: Declaring and using 8-bit integers in the `Private`,
  `Workgroup` (for non-Block variables), and `Function` storage classes
  is enabled, while declaring them in the interface storage classes
  (e.g., `UniformConstant`, `Uniform`, `StorageBuffer`, `Input`,
  `Output`, and `PushConstant`) is not enabled.

If the `VkPhysicalDeviceShaderFloat16Int8Features` structure is included
in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceShaderFloat16Int8Features` **can** also be used in the
`pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceShaderFloat16Int8Features-sType-sType"
  id="VUID-VkPhysicalDeviceShaderFloat16Int8Features-sType-sType"></a>
  VUID-VkPhysicalDeviceShaderFloat16Int8Features-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_FLOAT16_INT8_FEATURES`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_shader_float16_int8](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_shader_float16_int8.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html), [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceShaderFloat16Int8Features"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
