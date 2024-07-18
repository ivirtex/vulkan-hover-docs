# CoalescedInputCountAMDX(3) Manual Page

## Name

CoalescedInputCountAMDX - Number of inputs coalesced for a coalescing
node in a work graph



## <a href="#_description" class="anchor"></a>Description

`CoalescedInputCountAMDX`  
Decorating a variable with the `CoalescedInputCountAMDX` built-in
decoration will make that variable contain the number of node dispatches
that the implementation coalesced into the input for the current shader.
This variable will take a value in the range \[1, arraySize), where
arraySize is the maximum size of the input payload array for the shader.

Valid Usage

- <a href="#VUID-CoalescedInputCountAMDX-CoalescedInputCountAMDX-09172"
  id="VUID-CoalescedInputCountAMDX-CoalescedInputCountAMDX-09172"></a>
  VUID-CoalescedInputCountAMDX-CoalescedInputCountAMDX-09172  
  The variable decorated with `CoalescedInputCountAMDX` **must** be
  declared using the `Input` `Storage` `Class`

- <a href="#VUID-CoalescedInputCountAMDX-CoalescedInputCountAMDX-09173"
  id="VUID-CoalescedInputCountAMDX-CoalescedInputCountAMDX-09173"></a>
  VUID-CoalescedInputCountAMDX-CoalescedInputCountAMDX-09173  
  If a variable is decorated with `CoalescedInputCountAMDX`, the
  `CoalescingAMDX` execution mode **must** be declared

- <a href="#VUID-CoalescedInputCountAMDX-CoalescedInputCountAMDX-09174"
  id="VUID-CoalescedInputCountAMDX-CoalescedInputCountAMDX-09174"></a>
  VUID-CoalescedInputCountAMDX-CoalescedInputCountAMDX-09174  
  The variable decorated with `CoalescedInputCountAMDX` **must** be
  declared as a scalar 32-bit integer value

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#CoalescedInputCountAMDX"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
