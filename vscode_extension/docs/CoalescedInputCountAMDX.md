# CoalescedInputCountAMDX(3) Manual Page

## Name

CoalescedInputCountAMDX - Number of inputs coalesced for a coalescing node in a work graph



## [](#_description)Description

`CoalescedInputCountAMDX`

Decorating a variable with the `CoalescedInputCountAMDX` built-in decoration will make that variable contain the number of node dispatches that the implementation coalesced into the input for the current shader. This variable will take a value in the range [1, arraySize), where arraySize is the maximum size of the input payload array for the shader.

Valid Usage

- [](#VUID-CoalescedInputCountAMDX-CoalescedInputCountAMDX-09172)VUID-CoalescedInputCountAMDX-CoalescedInputCountAMDX-09172  
  The variable decorated with `CoalescedInputCountAMDX` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-CoalescedInputCountAMDX-CoalescedInputCountAMDX-09173)VUID-CoalescedInputCountAMDX-CoalescedInputCountAMDX-09173  
  If a variable is decorated with `CoalescedInputCountAMDX`, the `CoalescingAMDX` execution mode **must** be declared
- [](#VUID-CoalescedInputCountAMDX-CoalescedInputCountAMDX-09174)VUID-CoalescedInputCountAMDX-CoalescedInputCountAMDX-09174  
  The variable decorated with `CoalescedInputCountAMDX` **must** be declared as a scalar 32-bit integer value

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#CoalescedInputCountAMDX).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0