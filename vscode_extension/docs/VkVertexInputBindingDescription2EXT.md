# VkVertexInputBindingDescription2EXT(3) Manual Page

## Name

VkVertexInputBindingDescription2EXT - Structure specifying the extended vertex input binding description



## [](#_c_specification)C Specification

The `VkVertexInputBindingDescription2EXT` structure is defined as:

```c++
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

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `binding` is the binding number that this structure describes.
- `stride` is the byte stride between consecutive elements within the buffer.
- `inputRate` is a [VkVertexInputRate](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVertexInputRate.html) value specifying whether vertex attribute addressing is a function of the vertex index or of the instance index.
- `divisor` is the number of successive instances that will use the same value of the vertex attribute when instanced rendering is enabled. This member **can** be a value other than `1` if the [`vertexAttributeInstanceRateDivisor`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-vertexAttributeInstanceRateDivisor) feature is enabled. For example, if the divisor is N, the same vertex attribute will be applied to N successive instances before moving on to the next vertex attribute. The maximum value of `divisor` is implementation-dependent and can be queried using `VkPhysicalDeviceVertexAttributeDivisorPropertiesEXT`::`maxVertexAttribDivisor`. A value of `0` **can** be used for the divisor if the [`vertexAttributeInstanceRateZeroDivisor`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-vertexAttributeInstanceRateZeroDivisor) feature is enabled. In this case, the same vertex attribute will be applied to all instances.

## [](#_description)Description

Valid Usage

- [](#VUID-VkVertexInputBindingDescription2EXT-binding-04796)VUID-VkVertexInputBindingDescription2EXT-binding-04796  
  `binding` **must** be less than `VkPhysicalDeviceLimits`::`maxVertexInputBindings`
- [](#VUID-VkVertexInputBindingDescription2EXT-stride-04797)VUID-VkVertexInputBindingDescription2EXT-stride-04797  
  `stride` **must** be less than or equal to `VkPhysicalDeviceLimits`::`maxVertexInputBindingStride`
- [](#VUID-VkVertexInputBindingDescription2EXT-divisor-04798)VUID-VkVertexInputBindingDescription2EXT-divisor-04798  
  If the [`vertexAttributeInstanceRateZeroDivisor`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-vertexAttributeInstanceRateZeroDivisor) feature is not enabled, `divisor` **must** not be `0`
- [](#VUID-VkVertexInputBindingDescription2EXT-divisor-04799)VUID-VkVertexInputBindingDescription2EXT-divisor-04799  
  If the [`vertexAttributeInstanceRateDivisor`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-vertexAttributeInstanceRateDivisor) feature is not enabled, `divisor` **must** be `1`
- [](#VUID-VkVertexInputBindingDescription2EXT-divisor-06226)VUID-VkVertexInputBindingDescription2EXT-divisor-06226  
  `divisor` **must** be a value between `0` and `VkPhysicalDeviceVertexAttributeDivisorPropertiesEXT`::`maxVertexAttribDivisor`, inclusive
- [](#VUID-VkVertexInputBindingDescription2EXT-divisor-06227)VUID-VkVertexInputBindingDescription2EXT-divisor-06227  
  If `divisor` is not `1` then `inputRate` **must** be of type `VK_VERTEX_INPUT_RATE_INSTANCE`

Valid Usage (Implicit)

- [](#VUID-VkVertexInputBindingDescription2EXT-sType-sType)VUID-VkVertexInputBindingDescription2EXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VERTEX_INPUT_BINDING_DESCRIPTION_2_EXT`
- [](#VUID-VkVertexInputBindingDescription2EXT-inputRate-parameter)VUID-VkVertexInputBindingDescription2EXT-inputRate-parameter  
  `inputRate` **must** be a valid [VkVertexInputRate](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVertexInputRate.html) value

## [](#_see_also)See Also

[VK\_EXT\_shader\_object](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_shader_object.html), [VK\_EXT\_vertex\_input\_dynamic\_state](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_vertex_input_dynamic_state.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkVertexInputRate](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVertexInputRate.html), [vkCmdSetVertexInputEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetVertexInputEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVertexInputBindingDescription2EXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0