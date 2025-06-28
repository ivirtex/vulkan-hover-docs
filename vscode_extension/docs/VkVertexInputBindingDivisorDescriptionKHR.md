# VkVertexInputBindingDivisorDescription(3) Manual Page

## Name

VkVertexInputBindingDivisorDescription - Structure specifying a divisor used in instanced rendering



## [](#_c_specification)C Specification

The individual divisor values per binding are specified using the `VkVertexInputBindingDivisorDescription` structure which is defined as:

```c++
// Provided by VK_VERSION_1_4
typedef struct VkVertexInputBindingDivisorDescription {
    uint32_t    binding;
    uint32_t    divisor;
} VkVertexInputBindingDivisorDescription;
```

or the equivalent

```c++
// Provided by VK_KHR_vertex_attribute_divisor
typedef VkVertexInputBindingDivisorDescription VkVertexInputBindingDivisorDescriptionKHR;
```

or the equivalent

```c++
// Provided by VK_EXT_vertex_attribute_divisor
typedef VkVertexInputBindingDivisorDescription VkVertexInputBindingDivisorDescriptionEXT;
```

## [](#_members)Members

- `binding` is the binding number for which the divisor is specified.
- `divisor` is the number of successive instances that will use the same value of the vertex attribute when instanced rendering is enabled. For example, if the divisor is N, the same vertex attribute will be applied to N successive instances before moving on to the next vertex attribute. The maximum value of `divisor` is implementation-dependent and can be queried using [VkPhysicalDeviceVertexAttributeDivisorProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVertexAttributeDivisorProperties.html)::`maxVertexAttribDivisor`. A value of `0` **can** be used for the divisor if the [`vertexAttributeInstanceRateZeroDivisor`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-vertexAttributeInstanceRateZeroDivisor) feature is enabled. In this case, the same vertex attribute will be applied to all instances.

## [](#_description)Description

If this structure is not used to define a divisor value for an attribute, then the divisor has a logical default value of 1.

Valid Usage

- [](#VUID-VkVertexInputBindingDivisorDescription-binding-01869)VUID-VkVertexInputBindingDivisorDescription-binding-01869  
  `binding` **must** be less than [VkPhysicalDeviceLimits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceLimits.html)::`maxVertexInputBindings`
- [](#VUID-VkVertexInputBindingDivisorDescription-vertexAttributeInstanceRateZeroDivisor-02228)VUID-VkVertexInputBindingDivisorDescription-vertexAttributeInstanceRateZeroDivisor-02228  
  If the [`vertexAttributeInstanceRateZeroDivisor`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-vertexAttributeInstanceRateZeroDivisor) feature is not enabled, `divisor` **must** not be `0`
- [](#VUID-VkVertexInputBindingDivisorDescription-vertexAttributeInstanceRateDivisor-02229)VUID-VkVertexInputBindingDivisorDescription-vertexAttributeInstanceRateDivisor-02229  
  If the [`vertexAttributeInstanceRateDivisor`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-vertexAttributeInstanceRateDivisor) feature is not enabled, `divisor` **must** be `1`
- [](#VUID-VkVertexInputBindingDivisorDescription-divisor-01870)VUID-VkVertexInputBindingDivisorDescription-divisor-01870  
  `divisor` **must** be a value between `0` and [VkPhysicalDeviceVertexAttributeDivisorProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVertexAttributeDivisorProperties.html)::`maxVertexAttribDivisor`, inclusive
- [](#VUID-VkVertexInputBindingDivisorDescription-inputRate-01871)VUID-VkVertexInputBindingDivisorDescription-inputRate-01871  
  [VkVertexInputBindingDescription](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVertexInputBindingDescription.html)::`inputRate` **must** be of type `VK_VERTEX_INPUT_RATE_INSTANCE` for this `binding`

## [](#_see_also)See Also

[VK\_EXT\_vertex\_attribute\_divisor](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_vertex_attribute_divisor.html), [VK\_KHR\_vertex\_attribute\_divisor](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_vertex_attribute_divisor.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkPipelineVertexInputDivisorStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineVertexInputDivisorStateCreateInfo.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVertexInputBindingDivisorDescription)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0