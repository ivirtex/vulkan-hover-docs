# VkPhysicalDeviceShaderIntegerDotProductProperties(3) Manual Page

## Name

VkPhysicalDeviceShaderIntegerDotProductProperties - Structure containing
information about integer dot product support for a physical device



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceShaderIntegerDotProductProperties` structure is
defined as:

``` c
// Provided by VK_VERSION_1_3
typedef struct VkPhysicalDeviceShaderIntegerDotProductProperties {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           integerDotProduct8BitUnsignedAccelerated;
    VkBool32           integerDotProduct8BitSignedAccelerated;
    VkBool32           integerDotProduct8BitMixedSignednessAccelerated;
    VkBool32           integerDotProduct4x8BitPackedUnsignedAccelerated;
    VkBool32           integerDotProduct4x8BitPackedSignedAccelerated;
    VkBool32           integerDotProduct4x8BitPackedMixedSignednessAccelerated;
    VkBool32           integerDotProduct16BitUnsignedAccelerated;
    VkBool32           integerDotProduct16BitSignedAccelerated;
    VkBool32           integerDotProduct16BitMixedSignednessAccelerated;
    VkBool32           integerDotProduct32BitUnsignedAccelerated;
    VkBool32           integerDotProduct32BitSignedAccelerated;
    VkBool32           integerDotProduct32BitMixedSignednessAccelerated;
    VkBool32           integerDotProduct64BitUnsignedAccelerated;
    VkBool32           integerDotProduct64BitSignedAccelerated;
    VkBool32           integerDotProduct64BitMixedSignednessAccelerated;
    VkBool32           integerDotProductAccumulatingSaturating8BitUnsignedAccelerated;
    VkBool32           integerDotProductAccumulatingSaturating8BitSignedAccelerated;
    VkBool32           integerDotProductAccumulatingSaturating8BitMixedSignednessAccelerated;
    VkBool32           integerDotProductAccumulatingSaturating4x8BitPackedUnsignedAccelerated;
    VkBool32           integerDotProductAccumulatingSaturating4x8BitPackedSignedAccelerated;
    VkBool32           integerDotProductAccumulatingSaturating4x8BitPackedMixedSignednessAccelerated;
    VkBool32           integerDotProductAccumulatingSaturating16BitUnsignedAccelerated;
    VkBool32           integerDotProductAccumulatingSaturating16BitSignedAccelerated;
    VkBool32           integerDotProductAccumulatingSaturating16BitMixedSignednessAccelerated;
    VkBool32           integerDotProductAccumulatingSaturating32BitUnsignedAccelerated;
    VkBool32           integerDotProductAccumulatingSaturating32BitSignedAccelerated;
    VkBool32           integerDotProductAccumulatingSaturating32BitMixedSignednessAccelerated;
    VkBool32           integerDotProductAccumulatingSaturating64BitUnsignedAccelerated;
    VkBool32           integerDotProductAccumulatingSaturating64BitSignedAccelerated;
    VkBool32           integerDotProductAccumulatingSaturating64BitMixedSignednessAccelerated;
} VkPhysicalDeviceShaderIntegerDotProductProperties;
```

or the equivalent

``` c
// Provided by VK_KHR_shader_integer_dot_product
typedef VkPhysicalDeviceShaderIntegerDotProductProperties VkPhysicalDeviceShaderIntegerDotProductPropertiesKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

## <a href="#_description" class="anchor"></a>Description

- `integerDotProduct8BitUnsignedAccelerated` is a boolean that will be
  `VK_TRUE` if the support for 8-bit unsigned dot product operations
  using the `OpUDotKHR` SPIR-V instruction is accelerated <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#devsandqueues-integer-dot-product-accelerated"
  target="_blank" rel="noopener">as defined below</a>.

- `integerDotProduct8BitSignedAccelerated` is a boolean that will be
  `VK_TRUE` if the support for 8-bit signed dot product operations using
  the `OpSDotKHR` SPIR-V instruction is accelerated <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#devsandqueues-integer-dot-product-accelerated"
  target="_blank" rel="noopener">as defined below</a>.

- `integerDotProduct8BitMixedSignednessAccelerated` is a boolean that
  will be `VK_TRUE` if the support for 8-bit mixed signedness dot
  product operations using the `OpSUDotKHR` SPIR-V instruction is
  accelerated <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#devsandqueues-integer-dot-product-accelerated"
  target="_blank" rel="noopener">as defined below</a>.

- `integerDotProduct4x8BitPackedUnsignedAccelerated` is a boolean that
  will be `VK_TRUE` if the support for 8-bit unsigned dot product
  operations from operands packed into 32-bit integers using the
  `OpUDotKHR` SPIR-V instruction is accelerated <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#devsandqueues-integer-dot-product-accelerated"
  target="_blank" rel="noopener">as defined below</a>.

- `integerDotProduct4x8BitPackedSignedAccelerated` is a boolean that
  will be `VK_TRUE` if the support for 8-bit signed dot product
  operations from operands packed into 32-bit integers using the
  `OpSDotKHR` SPIR-V instruction is accelerated <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#devsandqueues-integer-dot-product-accelerated"
  target="_blank" rel="noopener">as defined below</a>.

- `integerDotProduct4x8BitPackedMixedSignednessAccelerated` is a boolean
  that will be `VK_TRUE` if the support for 8-bit mixed signedness dot
  product operations from operands packed into 32-bit integers using the
  `OpSUDotKHR` SPIR-V instruction is accelerated <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#devsandqueues-integer-dot-product-accelerated"
  target="_blank" rel="noopener">as defined below</a>.

- `integerDotProduct16BitUnsignedAccelerated` is a boolean that will be
  `VK_TRUE` if the support for 16-bit unsigned dot product operations
  using the `OpUDotKHR` SPIR-V instruction is accelerated <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#devsandqueues-integer-dot-product-accelerated"
  target="_blank" rel="noopener">as defined below</a>.

- `integerDotProduct16BitSignedAccelerated` is a boolean that will be
  `VK_TRUE` if the support for 16-bit signed dot product operations
  using the `OpSDotKHR` SPIR-V instruction is accelerated <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#devsandqueues-integer-dot-product-accelerated"
  target="_blank" rel="noopener">as defined below</a>.

- `integerDotProduct16BitMixedSignednessAccelerated` is a boolean that
  will be `VK_TRUE` if the support for 16-bit mixed signedness dot
  product operations using the `OpSUDotKHR` SPIR-V instruction is
  accelerated <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#devsandqueues-integer-dot-product-accelerated"
  target="_blank" rel="noopener">as defined below</a>.

- `integerDotProduct32BitUnsignedAccelerated` is a boolean that will be
  `VK_TRUE` if the support for 32-bit unsigned dot product operations
  using the `OpUDotKHR` SPIR-V instruction is accelerated <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#devsandqueues-integer-dot-product-accelerated"
  target="_blank" rel="noopener">as defined below</a>.

- `integerDotProduct32BitSignedAccelerated` is a boolean that will be
  `VK_TRUE` if the support for 32-bit signed dot product operations
  using the `OpSDotKHR` SPIR-V instruction is accelerated <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#devsandqueues-integer-dot-product-accelerated"
  target="_blank" rel="noopener">as defined below</a>.

- `integerDotProduct32BitMixedSignednessAccelerated` is a boolean that
  will be `VK_TRUE` if the support for 32-bit mixed signedness dot
  product operations using the `OpSUDotKHR` SPIR-V instruction is
  accelerated <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#devsandqueues-integer-dot-product-accelerated"
  target="_blank" rel="noopener">as defined below</a>.

- `integerDotProduct64BitUnsignedAccelerated` is a boolean that will be
  `VK_TRUE` if the support for 64-bit unsigned dot product operations
  using the `OpUDotKHR` SPIR-V instruction is accelerated <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#devsandqueues-integer-dot-product-accelerated"
  target="_blank" rel="noopener">as defined below</a>.

- `integerDotProduct64BitSignedAccelerated` is a boolean that will be
  `VK_TRUE` if the support for 64-bit signed dot product operations
  using the `OpSDotKHR` SPIR-V instruction is accelerated <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#devsandqueues-integer-dot-product-accelerated"
  target="_blank" rel="noopener">as defined below</a>.

- `integerDotProduct64BitMixedSignednessAccelerated` is a boolean that
  will be `VK_TRUE` if the support for 64-bit mixed signedness dot
  product operations using the `OpSUDotKHR` SPIR-V instruction is
  accelerated <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#devsandqueues-integer-dot-product-accelerated"
  target="_blank" rel="noopener">as defined below</a>.

- `integerDotProductAccumulatingSaturating8BitUnsignedAccelerated` is a
  boolean that will be `VK_TRUE` if the support for 8-bit unsigned
  accumulating saturating dot product operations using the
  `OpUDotAccSatKHR` SPIR-V instruction is accelerated <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#devsandqueues-integer-dot-product-accelerated"
  target="_blank" rel="noopener">as defined below</a>.

- `integerDotProductAccumulatingSaturating8BitSignedAccelerated` is a
  boolean that will be `VK_TRUE` if the support for 8-bit signed
  accumulating saturating dot product operations using the
  `OpSDotAccSatKHR` SPIR-V instruction is accelerated <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#devsandqueues-integer-dot-product-accelerated"
  target="_blank" rel="noopener">as defined below</a>.

- `integerDotProductAccumulatingSaturating8BitMixedSignednessAccelerated`
  is a boolean that will be `VK_TRUE` if the support for 8-bit mixed
  signedness accumulating saturating dot product operations using the
  `OpSUDotAccSatKHR` SPIR-V instruction is accelerated <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#devsandqueues-integer-dot-product-accelerated"
  target="_blank" rel="noopener">as defined below</a>.

- `integerDotProductAccumulatingSaturating4x8BitPackedUnsignedAccelerated`
  is a boolean that will be `VK_TRUE` if the support for 8-bit unsigned
  accumulating saturating dot product operations from operands packed
  into 32-bit integers using the `OpUDotAccSatKHR` SPIR-V instruction is
  accelerated <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#devsandqueues-integer-dot-product-accelerated"
  target="_blank" rel="noopener">as defined below</a>.

- `integerDotProductAccumulatingSaturating4x8BitPackedSignedAccelerated`
  is a boolean that will be `VK_TRUE` if the support for 8-bit signed
  accumulating saturating dot product operations from operands packed
  into 32-bit integers using the `OpSDotAccSatKHR` SPIR-V instruction is
  accelerated <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#devsandqueues-integer-dot-product-accelerated"
  target="_blank" rel="noopener">as defined below</a>.

- `integerDotProductAccumulatingSaturating4x8BitPackedMixedSignednessAccelerated`
  is a boolean that will be `VK_TRUE` if the support for 8-bit mixed
  signedness accumulating saturating dot product operations from
  operands packed into 32-bit integers using the `OpSUDotAccSatKHR`
  SPIR-V instruction is accelerated <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#devsandqueues-integer-dot-product-accelerated"
  target="_blank" rel="noopener">as defined below</a>.

- `integerDotProductAccumulatingSaturating16BitUnsignedAccelerated` is a
  boolean that will be `VK_TRUE` if the support for 16-bit unsigned
  accumulating saturating dot product operations using the
  `OpUDotAccSatKHR` SPIR-V instruction is accelerated <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#devsandqueues-integer-dot-product-accelerated"
  target="_blank" rel="noopener">as defined below</a>.

- `integerDotProductAccumulatingSaturating16BitSignedAccelerated` is a
  boolean that will be `VK_TRUE` if the support for 16-bit signed
  accumulating saturating dot product operations using the
  `OpSDotAccSatKHR` SPIR-V instruction is accelerated <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#devsandqueues-integer-dot-product-accelerated"
  target="_blank" rel="noopener">as defined below</a>.

- `integerDotProductAccumulatingSaturating16BitMixedSignednessAccelerated`
  is a boolean that will be `VK_TRUE` if the support for 16-bit mixed
  signedness accumulating saturating dot product operations using the
  `OpSUDotAccSatKHR` SPIR-V instruction is accelerated <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#devsandqueues-integer-dot-product-accelerated"
  target="_blank" rel="noopener">as defined below</a>.

- `integerDotProductAccumulatingSaturating32BitUnsignedAccelerated` is a
  boolean that will be `VK_TRUE` if the support for 32-bit unsigned
  accumulating saturating dot product operations using the
  `OpUDotAccSatKHR` SPIR-V instruction is accelerated <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#devsandqueues-integer-dot-product-accelerated"
  target="_blank" rel="noopener">as defined below</a>.

- `integerDotProductAccumulatingSaturating32BitSignedAccelerated` is a
  boolean that will be `VK_TRUE` if the support for 32-bit signed
  accumulating saturating dot product operations using the
  `OpSDotAccSatKHR` SPIR-V instruction is accelerated <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#devsandqueues-integer-dot-product-accelerated"
  target="_blank" rel="noopener">as defined below</a>.

- `integerDotProductAccumulatingSaturating32BitMixedSignednessAccelerated`
  is a boolean that will be `VK_TRUE` if the support for 32-bit mixed
  signedness accumulating saturating dot product operations using the
  `OpSUDotAccSatKHR` SPIR-V instruction is accelerated <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#devsandqueues-integer-dot-product-accelerated"
  target="_blank" rel="noopener">as defined below</a>.

- `integerDotProductAccumulatingSaturating64BitUnsignedAccelerated` is a
  boolean that will be `VK_TRUE` if the support for 64-bit unsigned
  accumulating saturating dot product operations using the
  `OpUDotAccSatKHR` SPIR-V instruction is accelerated <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#devsandqueues-integer-dot-product-accelerated"
  target="_blank" rel="noopener">as defined below</a>.

- `integerDotProductAccumulatingSaturating64BitSignedAccelerated` is a
  boolean that will be `VK_TRUE` if the support for 64-bit signed
  accumulating saturating dot product operations using the
  `OpSDotAccSatKHR` SPIR-V instruction is accelerated <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#devsandqueues-integer-dot-product-accelerated"
  target="_blank" rel="noopener">as defined below</a>.

- `integerDotProductAccumulatingSaturating64BitMixedSignednessAccelerated`
  is a boolean that will be `VK_TRUE` if the support for 64-bit mixed
  signedness accumulating saturating dot product operations using the
  `OpSUDotAccSatKHR` SPIR-V instruction is accelerated <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#devsandqueues-integer-dot-product-accelerated"
  target="_blank" rel="noopener">as defined below</a>.

If the `VkPhysicalDeviceShaderIntegerDotProductProperties` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

These are properties of the integer dot product acceleration information
of a physical device.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>A dot product operation is deemed accelerated if its implementation
provides a performance advantage over application-provided code composed
from elementary instructions and/or other dot product instructions,
either because the implementation uses optimized machine code sequences
whose generation from application-provided code cannot be guaranteed or
because it uses hardware features that cannot otherwise be targeted from
application-provided code.</p></td>
</tr>
</tbody>
</table>

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceShaderIntegerDotProductProperties-sType-sType"
  id="VUID-VkPhysicalDeviceShaderIntegerDotProductProperties-sType-sType"></a>
  VUID-VkPhysicalDeviceShaderIntegerDotProductProperties-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_INTEGER_DOT_PRODUCT_PROPERTIES`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_shader_integer_dot_product](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_shader_integer_dot_product.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html), [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceShaderIntegerDotProductProperties"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
