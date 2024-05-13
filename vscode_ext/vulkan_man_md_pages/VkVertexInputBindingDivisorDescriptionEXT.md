# VkVertexInputBindingDivisorDescriptionKHR(3) Manual Page

## Name

VkVertexInputBindingDivisorDescriptionKHR - Structure specifying a
divisor used in instanced rendering



## <a href="#_c_specification" class="anchor"></a>C Specification

The individual divisor values per binding are specified using the
`VkVertexInputBindingDivisorDescriptionKHR` structure which is defined
as:

``` c
// Provided by VK_KHR_vertex_attribute_divisor
typedef struct VkVertexInputBindingDivisorDescriptionKHR {
    uint32_t    binding;
    uint32_t    divisor;
} VkVertexInputBindingDivisorDescriptionKHR;
```

or the equivalent

``` c
// Provided by VK_EXT_vertex_attribute_divisor
typedef VkVertexInputBindingDivisorDescriptionKHR VkVertexInputBindingDivisorDescriptionEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `binding` is the binding number for which the divisor is specified.

- `divisor` is the number of successive instances that will use the same
  value of the vertex attribute when instanced rendering is enabled. For
  example, if the divisor is N, the same vertex attribute will be
  applied to N successive instances before moving on to the next vertex
  attribute. The maximum value of `divisor` is implementation-dependent
  and can be queried using
  [VkPhysicalDeviceVertexAttributeDivisorPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVertexAttributeDivisorPropertiesKHR.html)::`maxVertexAttribDivisor`.
  A value of `0` **can** be used for the divisor if the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-vertexAttributeInstanceRateZeroDivisor"
  target="_blank"
  rel="noopener"><code>vertexAttributeInstanceRateZeroDivisor</code></a>
  feature is enabled. In this case, the same vertex attribute will be
  applied to all instances.

## <a href="#_description" class="anchor"></a>Description

If this structure is not used to define a divisor value for an
attribute, then the divisor has a logical default value of 1.

Valid Usage

- <a href="#VUID-VkVertexInputBindingDivisorDescriptionKHR-binding-01869"
  id="VUID-VkVertexInputBindingDivisorDescriptionKHR-binding-01869"></a>
  VUID-VkVertexInputBindingDivisorDescriptionKHR-binding-01869  
  `binding` **must** be less than
  [VkPhysicalDeviceLimits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLimits.html)::`maxVertexInputBindings`

- <a
  href="#VUID-VkVertexInputBindingDivisorDescriptionKHR-vertexAttributeInstanceRateZeroDivisor-02228"
  id="VUID-VkVertexInputBindingDivisorDescriptionKHR-vertexAttributeInstanceRateZeroDivisor-02228"></a>
  VUID-VkVertexInputBindingDivisorDescriptionKHR-vertexAttributeInstanceRateZeroDivisor-02228  
  If the `vertexAttributeInstanceRateZeroDivisor` feature is not
  enabled, `divisor` **must** not be `0`

- <a
  href="#VUID-VkVertexInputBindingDivisorDescriptionKHR-vertexAttributeInstanceRateDivisor-02229"
  id="VUID-VkVertexInputBindingDivisorDescriptionKHR-vertexAttributeInstanceRateDivisor-02229"></a>
  VUID-VkVertexInputBindingDivisorDescriptionKHR-vertexAttributeInstanceRateDivisor-02229  
  If the `vertexAttributeInstanceRateDivisor` feature is not enabled,
  `divisor` **must** be `1`

- <a href="#VUID-VkVertexInputBindingDivisorDescriptionKHR-divisor-01870"
  id="VUID-VkVertexInputBindingDivisorDescriptionKHR-divisor-01870"></a>
  VUID-VkVertexInputBindingDivisorDescriptionKHR-divisor-01870  
  `divisor` **must** be a value between `0` and
  [VkPhysicalDeviceVertexAttributeDivisorPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVertexAttributeDivisorPropertiesKHR.html)::`maxVertexAttribDivisor`,
  inclusive

- <a
  href="#VUID-VkVertexInputBindingDivisorDescriptionKHR-inputRate-01871"
  id="VUID-VkVertexInputBindingDivisorDescriptionKHR-inputRate-01871"></a>
  VUID-VkVertexInputBindingDivisorDescriptionKHR-inputRate-01871  
  [VkVertexInputBindingDescription](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVertexInputBindingDescription.html)::`inputRate`
  **must** be of type `VK_VERTEX_INPUT_RATE_INSTANCE` for this `binding`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_vertex_attribute_divisor](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_vertex_attribute_divisor.html),
[VK_KHR_vertex_attribute_divisor](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_vertex_attribute_divisor.html),
[VkPipelineVertexInputDivisorStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineVertexInputDivisorStateCreateInfoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVertexInputBindingDivisorDescriptionKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
