# VkPhysicalDeviceShaderEnqueueFeaturesAMDX(3) Manual Page

## Name

VkPhysicalDeviceShaderEnqueueFeaturesAMDX - Structure describing whether
shader enqueue within execution graphs are supported by the
implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceShaderEnqueueFeaturesAMDX` structure is defined as:

``` c
// Provided by VK_AMDX_shader_enqueue
typedef struct VkPhysicalDeviceShaderEnqueueFeaturesAMDX {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           shaderEnqueue;
} VkPhysicalDeviceShaderEnqueueFeaturesAMDX;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following feature:

- <span id="features-shaderEnqueue"></span> `shaderEnqueue` indicates
  whether the implementation supports <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#executiongraphs"
  target="_blank" rel="noopener">execution graphs</a>.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceShaderEnqueueFeaturesAMDX` structure is included
in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceShaderEnqueueFeaturesAMDX` **can** also be used in the
`pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceShaderEnqueueFeaturesAMDX-sType-sType"
  id="VUID-VkPhysicalDeviceShaderEnqueueFeaturesAMDX-sType-sType"></a>
  VUID-VkPhysicalDeviceShaderEnqueueFeaturesAMDX-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_ENQUEUE_FEATURES_AMDX`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_AMDX_shader_enqueue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_AMDX_shader_enqueue.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceShaderEnqueueFeaturesAMDX"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
