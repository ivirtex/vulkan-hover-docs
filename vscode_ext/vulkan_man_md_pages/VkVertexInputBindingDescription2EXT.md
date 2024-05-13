# VkVertexInputBindingDescription2EXT(3) Manual Page

## Name

VkVertexInputBindingDescription2EXT - Structure specifying the extended
vertex input binding description



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkVertexInputBindingDescription2EXT` structure is defined as:

``` c
// Provided by VK_EXT_shader_object, VK_EXT_vertex_input_dynamic_state
typedef struct VkVertexInputBindingDescription2EXT {
    VkStructureType      sType;
    void*                pNext;
    uint32_t             binding;
    uint32_t             stride;
    VkVertexInputRate    inputRate;
    uint32_t             divisor;
} VkVertexInputBindingDescription2EXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `binding` is the binding number that this structure describes.

- `stride` is the byte stride between consecutive elements within the
  buffer.

- `inputRate` is a [VkVertexInputRate](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVertexInputRate.html) value
  specifying whether vertex attribute addressing is a function of the
  vertex index or of the instance index.

- `divisor` is the number of successive instances that will use the same
  value of the vertex attribute when instanced rendering is enabled.
  This member **can** be set to a value other than `1` if the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-vertexAttributeInstanceRateDivisor"
  target="_blank"
  rel="noopener"><code>vertexAttributeInstanceRateDivisor</code></a>
  feature is enabled. For example, if the divisor is N, the same vertex
  attribute will be applied to N successive instances before moving on
  to the next vertex attribute. The maximum value of `divisor` is
  implementation-dependent and can be queried using
  `VkPhysicalDeviceVertexAttributeDivisorPropertiesEXT`::`maxVertexAttribDivisor`.
  A value of `0` **can** be used for the divisor if the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-vertexAttributeInstanceRateZeroDivisor"
  target="_blank"
  rel="noopener"><code>vertexAttributeInstanceRateZeroDivisor</code></a>
  feature is enabled. In this case, the same vertex attribute will be
  applied to all instances.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkVertexInputBindingDescription2EXT-binding-04796"
  id="VUID-VkVertexInputBindingDescription2EXT-binding-04796"></a>
  VUID-VkVertexInputBindingDescription2EXT-binding-04796  
  `binding` **must** be less than
  `VkPhysicalDeviceLimits`::`maxVertexInputBindings`

- <a href="#VUID-VkVertexInputBindingDescription2EXT-stride-04797"
  id="VUID-VkVertexInputBindingDescription2EXT-stride-04797"></a>
  VUID-VkVertexInputBindingDescription2EXT-stride-04797  
  `stride` **must** be less than or equal to
  `VkPhysicalDeviceLimits`::`maxVertexInputBindingStride`

- <a href="#VUID-VkVertexInputBindingDescription2EXT-divisor-04798"
  id="VUID-VkVertexInputBindingDescription2EXT-divisor-04798"></a>
  VUID-VkVertexInputBindingDescription2EXT-divisor-04798  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-vertexAttributeInstanceRateZeroDivisor"
  target="_blank"
  rel="noopener"><code>vertexAttributeInstanceRateZeroDivisor</code></a>
  feature is not enabled, `divisor` **must** not be `0`

- <a href="#VUID-VkVertexInputBindingDescription2EXT-divisor-04799"
  id="VUID-VkVertexInputBindingDescription2EXT-divisor-04799"></a>
  VUID-VkVertexInputBindingDescription2EXT-divisor-04799  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-vertexAttributeInstanceRateDivisor"
  target="_blank"
  rel="noopener"><code>vertexAttributeInstanceRateDivisor</code></a>
  feature is not enabled, `divisor` **must** be `1`

- <a href="#VUID-VkVertexInputBindingDescription2EXT-divisor-06226"
  id="VUID-VkVertexInputBindingDescription2EXT-divisor-06226"></a>
  VUID-VkVertexInputBindingDescription2EXT-divisor-06226  
  `divisor` **must** be a value between `0` and
  `VkPhysicalDeviceVertexAttributeDivisorPropertiesEXT`::`maxVertexAttribDivisor`,
  inclusive

- <a href="#VUID-VkVertexInputBindingDescription2EXT-divisor-06227"
  id="VUID-VkVertexInputBindingDescription2EXT-divisor-06227"></a>
  VUID-VkVertexInputBindingDescription2EXT-divisor-06227  
  If `divisor` is not `1` then `inputRate` **must** be of type
  `VK_VERTEX_INPUT_RATE_INSTANCE`

Valid Usage (Implicit)

- <a href="#VUID-VkVertexInputBindingDescription2EXT-sType-sType"
  id="VUID-VkVertexInputBindingDescription2EXT-sType-sType"></a>
  VUID-VkVertexInputBindingDescription2EXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_VERTEX_INPUT_BINDING_DESCRIPTION_2_EXT`

- <a href="#VUID-VkVertexInputBindingDescription2EXT-inputRate-parameter"
  id="VUID-VkVertexInputBindingDescription2EXT-inputRate-parameter"></a>
  VUID-VkVertexInputBindingDescription2EXT-inputRate-parameter  
  `inputRate` **must** be a valid
  [VkVertexInputRate](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVertexInputRate.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_shader_object](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_shader_object.html),
[VK_EXT_vertex_input_dynamic_state](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_vertex_input_dynamic_state.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkVertexInputRate](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVertexInputRate.html),
[vkCmdSetVertexInputEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetVertexInputEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVertexInputBindingDescription2EXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
